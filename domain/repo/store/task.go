package store

import (
	"context"

	"github.com/viettranx/service-context/core"

	"github.com/quocdaitrn/cp-task/domain/entity"
)

// TaskRepo provides methods for interacting with task data.
type TaskRepo interface {
	// InsertOne inserts a task to database.
	InsertOne(ctx context.Context, task *entity.Task) error

	// UpdateOne updates a task to database.
	UpdateOne(ctx context.Context, task *entity.Task) error

	// DeleteOne deletes a task from database.
	DeleteOne(ctx context.Context, id uint) error

	// FindOne fetches a task from database by id.
	FindOne(ctx context.Context, id uint) (*entity.Task, error)

	// FindRangeByCriteria fetches list of tasks by criteria.
	FindRangeByCriteria(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Task, error)
}
