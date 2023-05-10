package storeimpl

import (
	"context"

	"github.com/pkg/errors"
	kiterrors "github.com/quocdaitrn/golang-kit/errors"
	"github.com/viettranx/service-context/core"
	"gorm.io/gorm"

	"github.com/quocdaitrn/cp-task/domain/entity"
	"github.com/quocdaitrn/cp-task/domain/repo/store"
)

// taskRepo implements methods of task's repository.
type taskRepo struct {
	db *gorm.DB
}

// NewTaskRepo creates and returns a new instances of UserRepo.
func NewTaskRepo(db *gorm.DB) store.TaskRepo {
	return &taskRepo{db: db}
}

// InsertOne inserts a task to database.
func (r *taskRepo) InsertOne(_ context.Context, task *entity.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// UpdateOne updates a task to database.
func (r *taskRepo) UpdateOne(_ context.Context, task *entity.Task) error {
	if err := r.db.Table(task.TableName()).Where("id = ?", task.ID).Updates(task).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// DeleteOne deletes a task from database.
func (r *taskRepo) DeleteOne(_ context.Context, id uint) error {
	// Soft delete
	if err := r.db.Table(entity.Task{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": entity.StatusDeleted,
		}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// FindOne fetches a task from database by id.
func (r *taskRepo) FindOne(_ context.Context, id uint) (*entity.Task, error) {
	var data entity.Task

	if err := r.db.
		Table(data.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, kiterrors.ErrRepoEntityNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &data, nil
}

// FindRangeByCriteria fetches list of tasks by criteria.
func (r *taskRepo) FindRangeByCriteria(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Task, error) {
	var tasks []entity.Task

	db := r.db.
		Table(entity.Task{}.TableName()).
		Where("status <> ?", entity.StatusDeleted)

	if filter.UserID != nil {
		db = db.Where("user_id = ?", *filter.UserID)
	}

	if filter.Status != nil {
		db = db.Where("status = ?", *filter.Status)
	}

	// Count total records match conditions
	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	// Query data with paging
	if err := db.Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&tasks).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return tasks, nil
}
