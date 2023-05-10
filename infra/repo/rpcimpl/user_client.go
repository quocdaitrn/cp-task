package rpcimpl

import (
	"context"

	"github.com/pkg/errors"
	kiterrors "github.com/quocdaitrn/golang-kit/errors"

	"github.com/quocdaitrn/cp-task/domain/entity"
	"github.com/quocdaitrn/cp-task/domain/repo/rpc"
	"github.com/quocdaitrn/cp-task/proto/pb"
)

type rpcClient struct {
	client pb.UserServiceClient
}

// NewUserRepo creates and returns a user repository to interact with user's domain.
func NewUserRepo(client pb.UserServiceClient) rpc.UserRepo {
	return &rpcClient{client: client}
}

// GetUsersByIDs gets list of users by list of ids.
func (c *rpcClient) GetUsersByIDs(ctx context.Context, ids []uint) ([]entity.SimpleUser, error) {
	userIDs := make([]int32, len(ids))

	for i := range ids {
		userIDs[i] = int32(ids[i])
	}

	resp, err := c.client.GetUsersByIDs(ctx, &pb.GetUsersByIDsRequest{Ids: userIDs})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	users := make([]entity.SimpleUser, len(resp.Users))

	for i := range users {
		respUser := resp.Users[i]
		users[i] = entity.SimpleUser{
			ID:        uint(respUser.Id),
			LastName:  respUser.LastName,
			FirstName: respUser.FirstName,
		}
	}

	return users, nil
}

// GetUserByID gets a user by id.
func (c *rpcClient) GetUserByID(ctx context.Context, id uint) (*entity.SimpleUser, error) {
	resp, err := c.client.GetUserByID(ctx, &pb.GetUserByIDRequest{Id: int32(id)})
	if err != nil {
		return nil, kiterrors.WithStack(err)
	}

	user := &entity.SimpleUser{
		ID:        uint(resp.User.Id),
		LastName:  resp.User.FirstName,
		FirstName: resp.User.LastName,
	}

	return user, nil
}
