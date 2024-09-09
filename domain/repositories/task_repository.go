package repositories

import (
	"context"

	"github.com/GraphZC/sdd-task-management/domain/models"
	"github.com/GraphZC/sdd-task-management/domain/requests"
)

type TaskRepository interface {
	Create(ctx context.Context, req *requests.TaskCreateRequest, userID string) (string, error)
	FindByID(ctx context.Context, taskID string) (*models.Task, error)
	FindByUserID(ctx context.Context, userID string) ([]models.Task, error)
	DeleteByID(ctx context.Context, taskID string) error
	UpdateByUD(ctx context.Context, taskID string, req *requests.TaskUpdateRequest) error
	UpdateStatusByID(ctx context.Context, taskID string, status string) error
}
