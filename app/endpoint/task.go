package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	golangkitauth "github.com/quocdaitrn/golang-kit/auth"

	"github.com/quocdaitrn/cp-task/domain/service"
)

// TaskServiceEndpoints is a set of domain service.TaskService's endpoints.
type TaskServiceEndpoints struct {
	ListTasksEndpoint     endpoint.Endpoint
	GetTaskEndpoint       endpoint.Endpoint
	CreateNewTaskEndpoint endpoint.Endpoint
	UpdateTaskEndpoint    endpoint.Endpoint
	DeleteTaskEndpoint    endpoint.Endpoint
}

// NewTaskServiceEndpoints creates and returns a new instance of
// TaskServiceEndpoints.
func NewTaskServiceEndpoints(
	svc service.TaskService,
	authClient golangkitauth.AuthenticateClient,
) *TaskServiceEndpoints {
	epts := &TaskServiceEndpoints{}

	epts.ListTasksEndpoint = newListTasksEndpoint(svc)
	epts.ListTasksEndpoint = golangkitauth.Authenticate(authClient)(epts.ListTasksEndpoint)

	epts.GetTaskEndpoint = newGetTaskEndpoint(svc)
	epts.GetTaskEndpoint = golangkitauth.Authenticate(authClient)(epts.GetTaskEndpoint)

	epts.CreateNewTaskEndpoint = newCreateNewTaskEndpoint(svc)
	epts.CreateNewTaskEndpoint = golangkitauth.Authenticate(authClient)(epts.CreateNewTaskEndpoint)

	epts.UpdateTaskEndpoint = newUpdateTaskEndpoint(svc)
	epts.UpdateTaskEndpoint = golangkitauth.Authenticate(authClient)(epts.UpdateTaskEndpoint)

	epts.DeleteTaskEndpoint = newDeleteTaskEndpoint(svc)
	epts.DeleteTaskEndpoint = golangkitauth.Authenticate(authClient)(epts.DeleteTaskEndpoint)

	return epts
}

// newListTasksEndpoint creates and returns a new endpoint for
// ListTasks use case.
func newListTasksEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.ListTasks(ctx, request.(*service.ListTasksRequest))
	}
}

// newGetTaskEndpoint creates and returns a new endpoint for
// GetTask use case.
func newGetTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.GetTask(ctx, request.(*service.GetTaskRequest))
	}
}

// newCreateNewTaskEndpoint creates and returns a new endpoint for
// CreateNewTask use case.
func newCreateNewTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.CreateNewTask(ctx, request.(*service.CreateNewTaskRequest))
	}
}

// newUpdateTaskEndpoint creates and returns a new endpoint for
// UpdateTask use case.
func newUpdateTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UpdateTask(ctx, request.(*service.UpdateTaskRequest))
	}
}

// newDeleteTaskEndpoint creates and returns a new endpoint for
// DeleteTask use case.
func newDeleteTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.DeleteTask(ctx, request.(*service.DeleteTaskRequest))
	}
}
