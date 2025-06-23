package grpc

import (
	"log"
	"microservices/pkg/proto-gen/userpb"
	"microservices/user/internal/service"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewGRPCServer(svc service.UserService) *Server {
	return &Server{
		userService: svc,
	}
}

func (s *Server) Start(port string) {

	go func() {
		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen on port %v: %v", port, err)
		}

		grpcServer := grpc.NewServer()
		userpb.RegisterUserServiceServer(grpcServer, s)

		log.Printf("gRPC server listening at %v", listener.Addr())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
