package service

import (
	"context"

	"github.com/quocdaitrn/cp-task/domain/entity"
)

// TaskService exposes all available use cases of task domain.
type TaskService interface {
	// CreateNewTask creates a new task.
	CreateNewTask(ctx context.Context, req *CreateNewTaskRequest) (*CreateNewTaskResponse, error)

	// GetTask finds and returns a specific task.
	GetTask(ctx context.Context, req *GetTaskRequest) (*GetTaskResponse, error)

	// ListTasks finds and returns a list of tasks.
	ListTasks(ctx context.Context, req *ListTasksRequest) (*ListTasksResponse, error)

	// UpdateTask updates a specific task.
	UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*UpdateTaskResponse, error)

	// DeleteTask deletes a specific task.
	DeleteTask(ctx context.Context, req *DeleteTaskRequest) (*DeleteTaskResponse, error)
}

// CreateNewTaskRequest represent a request to create a task.
type CreateNewTaskRequest struct {
	Title       string `json:"title" validate:"required,max=256"`
	Description string `json:"description" validate:"required"`
}

// CreateNewTaskResponse represent a response for creating a task.
type CreateNewTaskResponse struct {
	Message string `json:"message"`
}

// GetTaskRequest represent a request to get a task.
type GetTaskRequest struct {
	ID string `param:"id" validate:"required"`
}

// GetTaskResponse represent a response for getting a task.
type GetTaskResponse struct {
	*entity.Task
}

// ListTasksRequest represent a request to get a list of tasks.
type ListTasksRequest struct {
	UserID *uint   `json:"-" query:"user_id" field:"user_id"`
	Status *string `json:"-" query:"status" field:"status"`
	Page   int     `json:"-" query:"page" field:"page" validate:"gte=1"`
	Limit  int     `json:"-" query:"limit" field:"limit" validate:"gte=1"`
}

// ListTasksResponse represent a response for listing tasks.
type ListTasksResponse struct {
	Items   []entity.Task `json:"items"`
	HasNext bool          `json:"has_next"`
	Page    uint          `json:"page"`
	Limit   uint          `json:"limit"`
}

// UpdateTaskRequest represent a request to update a task.
type UpdateTaskRequest struct {
	ID          string  `json:"-"  param:"id" validate:"required"`
	Title       *string `json:"title" validate:"max=256"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

// UpdateTaskResponse represent a response for updating a task.
type UpdateTaskResponse struct {
	Message string `json:"message"`
}

// DeleteTaskRequest represent a request to delete a task.
type DeleteTaskRequest struct {
	ID string `param:"id" validate:"required"`
}

// DeleteTaskResponse represent a response for deleting a task.
type DeleteTaskResponse struct {
	Message string `json:"message"`
}
