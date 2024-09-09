package mysql

import (
	"context"
	"database/sql"

	"github.com/GraphZC/sdd-task-management/domain/models"
	"github.com/GraphZC/sdd-task-management/domain/repositories"
	"github.com/GraphZC/sdd-task-management/domain/requests"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserMySQLRepository struct {
	db *sqlx.DB
}

func NewUserMySQLRepository(db *sqlx.DB) repositories.UserRepository {
	return &UserMySQLRepository{
		db: db,
	}
}

func (u *UserMySQLRepository) Create(ctx context.Context, req *requests.UserRegisterRequest) error {
	// Generate UUID
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	_, err = u.db.QueryContext(ctx, "INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", id.String(), req.Name, req.Email, req.Password)

	return err
}

func (u *UserMySQLRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?", email)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
