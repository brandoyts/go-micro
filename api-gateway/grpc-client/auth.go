package grpcClient

import (
	"microservices/pkg/proto-gen/authpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAuthService(addr string) (authpb.AuthServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := authpb.NewAuthServiceClient(conn)

	return client, nil
}
