package rpc

import (
	"context"

	"github.com/quocdaitrn/cp-task/domain/entity"
)

// UserRepo provides all methods to interact with user's domain.
type UserRepo interface {
	// GetUsersByIDs gets list of users by list of ids.
	GetUsersByIDs(ctx context.Context, ids []uint) ([]entity.SimpleUser, error)

	// GetUserByID gets a user by id.
	GetUserByID(ctx context.Context, id uint) (*entity.SimpleUser, error)
}
