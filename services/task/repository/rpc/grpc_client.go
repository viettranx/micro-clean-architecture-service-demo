package rpc

import (
	"context"
	"demo-service/proto/pb"
	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
)

type rpcClient struct {
	client pb.UserServiceClient
}

func NewClient(client pb.UserServiceClient) *rpcClient {
	return &rpcClient{client: client}
}

func (c *rpcClient) GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error) {
	userIds := make([]int32, len(ids))

	for i := range ids {
		userIds[i] = int32(ids[i])
	}

	resp, err := c.client.GetUsersByIds(ctx, &pb.GetUsersByIdsReq{Ids: userIds})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	users := make([]core.SimpleUser, len(resp.Users))

	for i := range users {
		respUser := resp.Users[i]
		users[i] = core.NewSimpleUser(int(respUser.Id), respUser.FirstName, respUser.LastName, nil)
	}

	return users, nil
}

func (c *rpcClient) GetUserById(ctx context.Context, id int) (*core.SimpleUser, error) {
	resp, err := c.client.GetUserById(ctx, &pb.GetUserByIdReq{Id: int32(id)})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	user := core.NewSimpleUser(int(resp.User.Id), resp.User.FirstName, resp.User.LastName, nil)

	return &user, nil
}
