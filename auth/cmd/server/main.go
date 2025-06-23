package main

import (
	"log"
	"microservices/auth/internal/application"
	"microservices/auth/internal/infrastructure/jwt"
	"microservices/auth/internal/infrastructure/userClient"
	interfaces "microservices/auth/internal/interfaces/grpc"
	"microservices/pkg/proto-gen/authpb"
	"net"
	"os"

	"google.golang.org/grpc"
)

func buildAuthHandler() *interfaces.AuthHandler {
	userClientAddress := os.Getenv("USER_SERVICE_ADDRESS")
	userClient := userClient.NewUserClient(userClientAddress)
	jwtService := jwt.NewJwtService("secret33210")
	loginUseCase := application.NewLoginUsecase(userClient, jwtService)

	return interfaces.NewAuthHandler(loginUseCase)
}

func main() {
	const port = ":50055"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", port, err)
	}

	authHandler := buildAuthHandler()

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, authHandler)

	log.Printf("gRPC server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
