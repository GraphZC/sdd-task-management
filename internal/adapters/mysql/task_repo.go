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

type TaskMySQLRepository struct {
	db *sqlx.DB
}

func NewTaskMySQLRepository(db *sqlx.DB) repositories.TaskRepository {
	return &TaskMySQLRepository{
		db: db,
	}
}

func (t *TaskMySQLRepository) Create(ctx context.Context, req *requests.TaskCreateRequest, userID string) (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	_, err = t.db.ExecContext(ctx, "INSERT INTO tasks (id, user_id, title, description, status, priority) VALUES (?, ?, ?, ?, ?, ?)", id.String(), userID, req.Title, req.Description, models.TaskStatusTodo, req.Priority)
	if err != nil {
		return "", err
	}

	return id.String(), err
}

func (t *TaskMySQLRepository) FindByID(ctx context.Context, taskID string) (*models.Task, error) {
	var task models.Task
	err := t.db.GetContext(ctx, &task, "SELECT id, user_id, title, description, status, priority, created_at, updated_at FROM tasks WHERE id = ?", taskID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &task, err
}

func (t *TaskMySQLRepository) FindByUserID(ctx context.Context, userID string) ([]models.Task, error) {
	var tasks []models.Task
	err := t.db.SelectContext(ctx, &tasks, "SELECT id, user_id, title, description, status, priority, created_at, updated_at FROM tasks WHERE user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskMySQLRepository) DeleteByID(ctx context.Context, taskID string) error {
	_, err := t.db.ExecContext(ctx, "DELETE FROM tasks WHERE id = ?", taskID)

	return err
}

func (t *TaskMySQLRepository) UpdateByUD(ctx context.Context, taskID string, req *requests.TaskCreateRequest) error {
	_, err := t.db.ExecContext(ctx, "UPDATE tasks SET title = ?, description = ?, priority = ? WHERE id = ?", req.Title, req.Description, req.Priority, taskID)

	return err
}

func (t *TaskMySQLRepository) UpdateStatusByID(ctx context.Context, taskID string, status string) error {
	_, err := t.db.ExecContext(ctx, "UPDATE tasks SET status = ? WHERE id = ?", status, taskID)

	return err
}
