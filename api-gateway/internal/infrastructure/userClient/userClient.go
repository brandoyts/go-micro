package userClient

import (
	"context"
	"fmt"
	"log"
	"microservices/pkg/proto-gen/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	GetUser(ctx context.Context, in *userpb.GetUserRequest) (*userpb.GetUserResponse, error)
	CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error)
}

type client struct {
	client userpb.UserServiceClient
}

func New(address string) Client {
	connection, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect to user service at %s: %v", address, err)
	}

	return &client{
		client: userpb.NewUserServiceClient(connection),
	}
}

func (uc *client) GetUser(ctx context.Context, in *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	response, err := uc.client.GetUser(ctx, in)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func (uc *client) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	response, err := uc.client.CreateUser(ctx, in)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}
