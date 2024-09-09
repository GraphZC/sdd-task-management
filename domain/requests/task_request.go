package requests

type TaskCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Priority    int    `json:"priority" validate:"required"`
}

type TaskUpdateRequest = TaskCreateRequest

type TaskUpdateStatusRequest struct {
	Status string `json:"status" validate:"required"`
}
