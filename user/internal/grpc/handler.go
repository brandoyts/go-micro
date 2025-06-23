package grpc

import (
	"context"
	"microservices/pkg/entities"
	"microservices/pkg/proto-gen/userpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetUser(ctx context.Context, in *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {

	result, err := s.userService.FindByUuid(ctx, in.Uuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:      result.Uuid,
			Username:  result.Username,
			CreatedAt: timestamppb.New(result.CreatedAt),
		},
	}, nil
}

func (s *Server) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	payload := entities.User{
		Username: in.Username,
		Name:     in.Name,
		Password: in.Password,
	}

	result, err := s.userService.CreateUser(ctx, &payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Uuid:      result.Uuid,
			Username:  result.Username,
			Name:      result.Name,
			CreatedAt: timestamppb.New(result.CreatedAt),
		},
	}, nil
}

func (s *Server) GetUserByUsername(ctx context.Context, in *userpb.GetUserByUsernameRequest) (*userpb.GetUserByUsernameResponse, error) {
	payload := entities.User{
		Username: in.Username,
	}

	result, err := s.userService.Find(ctx, &payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &userpb.GetUserByUsernameResponse{
		User: &userpb.User{
			Uuid:      result.Uuid,
			Username:  result.Username,
			Password:  result.Password,
			CreatedAt: timestamppb.New(result.CreatedAt),
		},
	}, nil
}
