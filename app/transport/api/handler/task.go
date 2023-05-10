package handler

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/quocdaitrn/golang-kit/auth"
	golangkithttp "github.com/quocdaitrn/golang-kit/http"

	"github.com/quocdaitrn/cp-task/app/endpoint"
	"github.com/quocdaitrn/cp-task/app/transport/api/codec"
	"github.com/quocdaitrn/cp-task/domain/service"
)

// MakeTaskHTTPHandler provides all task's routes.
func MakeTaskHTTPHandler(
	r *mux.Router,
	svc service.TaskService,
	logger log.Logger,
	authClient auth.AuthenticateClient,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(golangkithttp.DefaultErrorEncoder),
		kithttp.ServerBefore(golangkithttp.PopulateRequestAuthorizationToken),
	}

	taskSvcEpts := endpoint.NewTaskServiceEndpoints(svc, authClient)

	getTaskHandler := kithttp.NewServer(
		taskSvcEpts.GetTaskEndpoint,
		codec.DecodeGetTaskRequest,
		golangkithttp.EncodeResponse,
		opts...,
	)

	listTasksHandler := kithttp.NewServer(
		taskSvcEpts.ListTasksEndpoint,
		codec.DecodeListTasksRequest,
		golangkithttp.EncodeResponse,
		opts...,
	)

	createTaskHandler := kithttp.NewServer(
		taskSvcEpts.CreateNewTaskEndpoint,
		codec.DecodeCreateTaskRequest,
		golangkithttp.EncodeResponse,
		opts...,
	)

	updateTaskHandler := kithttp.NewServer(
		taskSvcEpts.UpdateTaskEndpoint,
		codec.DecodeUpdateTaskRequest,
		golangkithttp.EncodeResponse,
		opts...,
	)

	deleteTaskHandler := kithttp.NewServer(
		taskSvcEpts.DeleteTaskEndpoint,
		codec.DecodeDeleteTaskRequest,
		golangkithttp.EncodeResponse,
		opts...,
	)

	r.Handle("/tasks/{id}", getTaskHandler).Methods(http.MethodGet)
	r.Handle("/tasks", listTasksHandler).Methods(http.MethodGet)
	r.Handle("/tasks", createTaskHandler).Methods(http.MethodPost)
	r.Handle("/tasks/{id}", updateTaskHandler).Methods(http.MethodPatch)
	r.Handle("/tasks/{id}", deleteTaskHandler).Methods(http.MethodDelete)

	return r
}
