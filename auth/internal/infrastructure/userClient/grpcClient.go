package userClient

import (
	"context"
	"errors"
	"log"
	"microservices/pkg/proto-gen/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	Uuid           string
	HashedPassword string
}

type UserClient interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type grpcUserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(addr string) UserClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service at %s: %v", addr, err)
	}

	return &grpcUserClient{
		client: userpb.NewUserServiceClient(conn),
	}
}

func (uc *grpcUserClient) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	req := &userpb.GetUserByUsernameRequest{
		Username: username,
	}

	resp, err := uc.client.GetUserByUsername(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp == nil || resp.User == nil {
		return nil, errors.New("user not found")
	}

	return &User{
		Uuid:           resp.User.Uuid,
		HashedPassword: resp.User.Password,
	}, nil
}
