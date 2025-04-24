package client

import (
	"context"
	"github.com/chocological13/tech-stream/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client user.UserServiceClient
}

func NewUserClient(address string) (*UserClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := user.NewUserServiceClient(conn)
	return &UserClient{
		client: client,
	}, nil
}

func (c *UserClient) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	return c.client.CreateUser(ctx, req)
}
