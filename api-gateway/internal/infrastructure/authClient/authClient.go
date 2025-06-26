package authClient

import (
	"context"
	"fmt"
	"log"
	"microservices/pkg/proto-gen/authpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error)
}

type client struct {
	client authpb.AuthServiceClient
}

func New(address string) Client {
	connection, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect to user service at %s: %v", address, err)
	}

	return &client{
		client: authpb.NewAuthServiceClient(connection),
	}
}

func (ac *client) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	response, err := ac.client.Login(ctx, in)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}
