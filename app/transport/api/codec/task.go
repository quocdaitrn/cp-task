package codec

import (
	"context"
	"net/http"

	kithttp "github.com/quocdaitrn/golang-kit/http"

	"github.com/quocdaitrn/cp-task/domain/service"
)

// DecodeGetTaskRequest decodes GetTaskRequest from http.Request.
func DecodeGetTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &service.GetTaskRequest{}
	if err := kithttp.Bind(r, req); err != nil {
		return nil, err
	}
	return req, nil
}

// DecodeListTasksRequest decodes ListTasksRequest from http.Request.
func DecodeListTasksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &service.ListTasksRequest{Limit: 20}
	if err := kithttp.Bind(r, req); err != nil {
		return nil, err
	}
	return req, nil
}

// DecodeCreateTaskRequest decodes CreateNewTaskRequest from http.Request.
func DecodeCreateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &service.CreateNewTaskRequest{}
	if err := kithttp.Bind(r, req); err != nil {
		return nil, err
	}
	return req, nil
}

// DecodeUpdateTaskRequest decodes UpdateTaskRequest from http.Request.
func DecodeUpdateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &service.UpdateTaskRequest{}
	if err := kithttp.Bind(r, req); err != nil {
		return nil, err
	}
	return req, nil
}

// DecodeDeleteTaskRequest decodes DeleteTaskRequest from http.Request.
func DecodeDeleteTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &service.DeleteTaskRequest{}
	if err := kithttp.Bind(r, req); err != nil {
		return nil, err
	}
	return req, nil
}
