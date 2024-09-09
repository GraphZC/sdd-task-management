package repositories

import (
	"context"

	"github.com/GraphZC/sdd-task-management/domain/models"
	"github.com/GraphZC/sdd-task-management/domain/requests"
)

type UserRepository interface {
	Create(ctx context.Context, req *requests.UserRegisterRequest) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
