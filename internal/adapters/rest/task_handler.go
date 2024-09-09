package rest

import (
	"github.com/GraphZC/sdd-task-management/domain/exceptions"
	"github.com/GraphZC/sdd-task-management/domain/requests"
	"github.com/GraphZC/sdd-task-management/domain/usecases"
	"github.com/GraphZC/sdd-task-management/utils"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler interface {
	CreateTask(c *fiber.Ctx) error
	FindTaskByID(c *fiber.Ctx) error
	FindTaskByUserID(c *fiber.Ctx) error
	DeleteTaskByID(c *fiber.Ctx) error
	UpdateTaskByID(c *fiber.Ctx) error
	UpdateTaskStatusByID(c *fiber.Ctx) error
}

type taskHandler struct {
	service usecases.TaskUseCase
}

func NewTaskHandler(service usecases.TaskUseCase) TaskHandler {
	return &taskHandler{
		service: service,
	}
}

func (t *taskHandler) CreateTask(c *fiber.Ctx) error {
	// Parse request
	var req *requests.TaskCreateRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Find id from jwt
	userID := utils.GetUserIDFromJWT(c)

	// Create task
	task, err := t.service.CreateTask(c.Context(), req, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (t *taskHandler) FindTaskByID(c *fiber.Ctx) error {
	// Get task ID
	taskID := c.Params("taskID")

	// Find id from jwt
	userID := utils.GetUserIDFromJWT(c)

	// Get task
	task, err := t.service.FindTaskByID(c.Context(), taskID, userID)
	if err != nil {
		switch err {
		case exceptions.ErrTaskNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Task not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func (t *taskHandler) FindTaskByUserID(c *fiber.Ctx) error {
	// Find id from jwt
	userID := utils.GetUserIDFromJWT(c)

	// Get tasks
	tasks, err := t.service.FindTaskByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(tasks)
}

func (t *taskHandler) DeleteTaskByID(c *fiber.Ctx) error {
	// Get task ID
	taskID := c.Params("taskID")

	// Find id from jwt
	userID := utils.GetUserIDFromJWT(c)

	// Delete task
	task, err := t.service.DeleteTaskByID(c.Context(), taskID, userID)
	if err != nil {
		switch err {
		case exceptions.ErrTaskNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Task not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func (t *taskHandler) UpdateTaskByID(c *fiber.Ctx) error {
	// Get task ID
	taskID := c.Params("taskID")

	// Parse request
	var req *requests.TaskUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Find id from jwt
	userID := utils.GetUserIDFromJWT(c)

	// Update task
	task, err := t.service.UpdateTaskByID(c.Context(), taskID, req, userID)
	if err != nil {
		switch err {
		case exceptions.ErrTaskNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Task not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func (t *taskHandler) UpdateTaskStatusByID(c *fiber.Ctx) error {
	// Get task ID
	taskID := c.Params("taskID")

	// Parse request
	var req *requests.TaskUpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Find id from jwt
	userID := utils.GetUserIDFromJWT(c)

	// Update task status
	task, err := t.service.UpdateTaskStatusByID(c.Context(), taskID, req, userID)
	if err != nil {
		switch err {
		case exceptions.ErrTaskNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Task not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(task)
}
