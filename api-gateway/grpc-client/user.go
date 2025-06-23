package grpcClient

import (
	"microservices/pkg/proto-gen/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitUserService(addr string) (userpb.UserServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := userpb.NewUserServiceClient(conn)

	return client, nil
}
