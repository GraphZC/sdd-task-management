package usecases

import (
	"context"

	"github.com/GraphZC/sdd-task-management/domain/exceptions"
	"github.com/GraphZC/sdd-task-management/domain/models"
	"github.com/GraphZC/sdd-task-management/domain/repositories"
	"github.com/GraphZC/sdd-task-management/domain/requests"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, req *requests.TaskCreateRequest, userID string) (*models.Task, error)
	FindTaskByID(ctx context.Context, taskID string, userID string) (*models.Task, error)
	FindTaskByUserID(ctx context.Context, userID string) ([]models.Task, error)
	DeleteTaskByID(ctx context.Context, taskID string, userID string) (*models.Task, error)
	UpdateTaskByID(ctx context.Context, taskID string, req *requests.TaskUpdateRequest, userID string) (*models.Task, error)
	UpdateTaskStatusByID(ctx context.Context, taskID string, req *requests.TaskUpdateStatusRequest, userID string) (*models.Task, error)
}

type taskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) TaskUseCase {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (t *taskService) CreateTask(ctx context.Context, req *requests.TaskCreateRequest, userID string) (*models.Task, error) {
	// Check priority
	if req.Priority < 0 || req.Priority > 2 {
		return nil, exceptions.ErrInvalidPriority
	}

	// Create task
	taskID, err := t.taskRepo.Create(ctx, req, userID)
	if err != nil {
		return nil, err
	}

	// Find the task
	task, err := t.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskService) FindTaskByID(ctx context.Context, taskID string, userID string) (*models.Task, error) {
	// Find the task
	task, err := t.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	// Check task is exist
	if task == nil {
		return nil, exceptions.ErrTaskNotFound
	}

	// Check task is belong to the user
	if task.UserID != userID {
		return nil, exceptions.ErrTaskNotFound
	}

	return task, nil
}

func (t *taskService) FindTaskByUserID(ctx context.Context, userID string) ([]models.Task, error) {
	return t.taskRepo.FindByUserID(ctx, userID)
}

func (t *taskService) DeleteTaskByID(ctx context.Context, taskID string, userID string) (*models.Task, error) {
	// Find the task
	task, err := t.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	// Check task is exist
	if task == nil {
		return nil, exceptions.ErrTaskNotFound
	}

	// Check task is belong to the user
	if task.UserID != userID {
		return nil, exceptions.ErrTaskNotFound
	}

	// Delete task in database
	err = t.taskRepo.DeleteByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskService) UpdateTaskByID(ctx context.Context, taskID string, req *requests.TaskCreateRequest, userID string) (*models.Task, error) {
	// Check priority
	if req.Priority < 0 || req.Priority > 2 {
		return nil, exceptions.ErrInvalidPriority
	}

	// Find the task
	task, err := t.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	// Check task is exist
	if task == nil {
		return nil, exceptions.ErrTaskNotFound
	}

	// Check task is belong to the user
	if task.UserID != userID {
		return nil, exceptions.ErrTaskNotFound
	}

	// Update task in database
	err = t.taskRepo.UpdateByUD(ctx, taskID, req)
	if err != nil {
		return nil, err
	}

	// Update task
	task.Title = req.Title
	task.Description = req.Description
	task.Priority = req.Priority

	return task, nil
}

func (t *taskService) UpdateTaskStatusByID(ctx context.Context, taskID string, req *requests.TaskUpdateStatusRequest, userID string) (*models.Task, error) {
	// Check status
	if req.Status != models.TaskStatusTodo && req.Status != models.TaskStatusCompleted {
		return nil, exceptions.ErrInvalidStatus
	}

	// Find the task
	task, err := t.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	// Check task is exist
	if task == nil {
		return nil, exceptions.ErrTaskNotFound
	}

	// Check task is belong to the user
	if task.UserID != userID {
		return nil, exceptions.ErrTaskNotFound
	}

	// Update status in database
	err = t.taskRepo.UpdateStatusByID(ctx, taskID, req.Status)
	if err != nil {
		return nil, err
	}

	// Update task
	task.Status = req.Status

	return task, nil
}
