package usecases

import (
	"context"
	"time"

	"github.com/GraphZC/sdd-task-management/configs"
	"github.com/GraphZC/sdd-task-management/domain/exceptions"
	"github.com/GraphZC/sdd-task-management/domain/repositories"
	"github.com/GraphZC/sdd-task-management/domain/requests"
	"github.com/GraphZC/sdd-task-management/domain/responses"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(ctx context.Context, req *requests.UserRegisterRequest) error
	Login(ctx context.Context, req *requests.UserLoginRequest) (*responses.UserLoginResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
	config   *configs.Config
}

func NewUserService(userRepo repositories.UserRepository, config *configs.Config) UserUseCase {
	return &userService{
		userRepo: userRepo,
		config:   config,
	}
}

func (u *userService) Register(ctx context.Context, req *requests.UserRegisterRequest) error {
	// Find user by email
	user, err := u.userRepo.FindByEmail(ctx, req.Email)

	if err != nil {
		return err
	}

	// Check if user already exist
	if user != nil {
		return exceptions.ErrDuplicatedEmail
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)

	return u.userRepo.Create(ctx, req)
}

func (u *userService) Login(ctx context.Context, req *requests.UserLoginRequest) (*responses.UserLoginResponse, error) {
	// Find user by email
	user, err := u.userRepo.FindByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	// Check if user exist
	if user == nil {
		return nil, exceptions.ErrLoginFailed
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, exceptions.ErrLoginFailed
	}

	// Generate JWT token
	expireAt := time.Now().Add(time.Hour * 1)

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   expireAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(u.config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &responses.UserLoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     tokenString,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
