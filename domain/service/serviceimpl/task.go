package serviceimpl

import (
	"context"

	kitcontext "github.com/quocdaitrn/golang-kit/context"
	kiterrors "github.com/quocdaitrn/golang-kit/errors"
	"github.com/quocdaitrn/golang-kit/validator"
	"github.com/viettranx/service-context/core"

	"github.com/quocdaitrn/cp-task/domain/entity"
	"github.com/quocdaitrn/cp-task/domain/repo/rpc"
	"github.com/quocdaitrn/cp-task/domain/repo/store"
	"github.com/quocdaitrn/cp-task/domain/service"
)

type taskService struct {
	taskRepo  store.TaskRepo
	userRepo  rpc.UserRepo
	validator validator.Validator
}

// NewTaskService creates and returns a new instance of TaskService.
func NewTaskService(
	taskRepo store.TaskRepo,
	userRepo rpc.UserRepo,
	validator validator.Validator,
) service.TaskService {
	return &taskService{
		taskRepo:  taskRepo,
		userRepo:  userRepo,
		validator: validator,
	}
}

// CreateNewTask creates a new task.
func (s *taskService) CreateNewTask(ctx context.Context, req *service.CreateNewTaskRequest) (*service.CreateNewTaskResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, kiterrors.WithStack(err)
	}

	uid := kitcontext.UIDFromContext(ctx)
	id, _ := core.FromBase58(uid.Sub)
	requesterID := uint(id.GetLocalID())

	task := &entity.Task{
		UserID:      requesterID,
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.StatusDoing,
	}
	if err := s.taskRepo.InsertOne(ctx, task); err != nil {
		return nil, err
	}

	return &service.CreateNewTaskResponse{Message: "create task successfully"}, nil
}

// GetTask finds and returns a specific task.
func (s *taskService) GetTask(ctx context.Context, req *service.GetTaskRequest) (*service.GetTaskResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, kiterrors.WithStack(err)
	}

	cUID, err := core.FromBase58(req.ID)

	task, err := s.taskRepo.FindOne(ctx, uint(cUID.GetLocalID()))
	if err != nil {
		if err == kiterrors.ErrRepoCacheNotFound {
			return nil, kiterrors.ErrNotFound
		}
		return nil, err
	}

	if task.Status == entity.StatusDeleted {
		return nil, kiterrors.ErrNotFound
	}

	// Get extra infos: User
	user, err := s.userRepo.GetUserByID(ctx, task.UserID)
	if err != nil {
		return nil, err
	}
	task.User = user
	task.Mask()

	return &service.GetTaskResponse{Task: task}, nil
}

// ListTasks finds and returns a list of tasks.
func (s *taskService) ListTasks(ctx context.Context, req *service.ListTasksRequest) (*service.ListTasksResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, kiterrors.WithStack(err)
	}

	paging := &core.Paging{
		Page:  req.Page,
		Limit: req.Limit,
	}
	tasks, err := s.taskRepo.FindRangeByCriteria(ctx,
		&entity.Filter{
			UserID: req.UserID,
			Status: req.Status,
		},
		paging,
	)
	if err != nil {
		return nil, err
	}

	// Get extra infos: User
	userIDs := make([]uint, len(tasks))

	for i := range userIDs {
		userIDs[i] = tasks[i].UserID
	}

	users, err := s.userRepo.GetUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	// For speed up mapping data
	userMap := make(map[uint]*entity.SimpleUser)

	for i, u := range users {
		userMap[u.ID] = &users[i]
	}

	for i, t := range tasks {
		tasks[i].User = userMap[t.UserID]
		tasks[i].Mask()
	}

	return &service.ListTasksResponse{
		Items:   tasks,
		HasNext: paging.Total > int64(req.Page*req.Limit),
		Page:    uint(req.Page),
		Limit:   uint(req.Limit),
	}, nil
}

// UpdateTask updates a specific task.
func (s *taskService) UpdateTask(ctx context.Context, req *service.UpdateTaskRequest) (*service.UpdateTaskResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, kiterrors.WithStack(err)
	}

	cUID, err := core.FromBase58(req.ID)

	// Get task data, without extra infos
	task, err := s.taskRepo.FindOne(ctx, uint(cUID.GetLocalID()))
	if err != nil {
		if err == kiterrors.ErrRepoEntityNotFound {
			return nil, kiterrors.ErrNotFound
		}

		return nil, err
	}

	uid := kitcontext.UIDFromContext(ctx)
	id, _ := core.FromBase58(uid.Sub)
	requesterID := uint(id.GetLocalID())

	// Only task owner can do this
	if requesterID != task.UserID {
		return nil, kiterrors.ErrForbidden.WithDetails("only owner can update their task")
	}

	// Only update task with doing status
	if task.Status != entity.StatusDoing {
		return nil, kiterrors.ErrForbidden.WithDetails("only update task with doing status")
	}

	if req.Title != nil && *req.Title != "" {
		task.Title = *req.Title
	}

	if req.Description != nil && *req.Description != "" {
		task.Description = *req.Description
	}

	if req.Status != nil && *req.Status != "" {
		task.Status = entity.Status(*req.Status)
	}

	if err := s.taskRepo.UpdateOne(ctx, task); err != nil {
		return nil, err
	}

	return &service.UpdateTaskResponse{Message: "update task successfully"}, nil
}

// DeleteTask deletes a specific task.
func (s *taskService) DeleteTask(ctx context.Context, req *service.DeleteTaskRequest) (*service.DeleteTaskResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, kiterrors.WithStack(err)
	}

	cUID, err := core.FromBase58(req.ID)
	localID := uint(cUID.GetLocalID())

	// Get task data, without extra infos
	task, err := s.taskRepo.FindOne(ctx, localID)
	if err != nil {
		if err == kiterrors.ErrRepoEntityNotFound {
			return nil, kiterrors.ErrNotFound
		}

		return nil, err
	}

	uid := kitcontext.UIDFromContext(ctx)
	id, _ := core.FromBase58(uid.Sub)
	requesterID := uint(id.GetLocalID())

	// Only task owner can do this
	if requesterID != task.UserID {
		return nil, kiterrors.ErrForbidden.WithDetails("only owner can delete their task")
	}

	// Only delete task with doing status
	if task.Status == entity.StatusDeleted {
		return nil, kiterrors.ErrForbidden.WithDetails("task has already deleted")
	}

	if err := s.taskRepo.DeleteOne(ctx, localID); err != nil {
		return nil, err
	}

	return &service.DeleteTaskResponse{Message: "delete task successfully"}, nil
}
