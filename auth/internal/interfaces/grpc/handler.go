package interfaces

import (
	"context"
	"microservices/auth/internal/application"
	"microservices/pkg/proto-gen/authpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	loginUseCase *application.LoginUsecase
}

func NewAuthHandler(loginUsecase *application.LoginUsecase) *AuthHandler {
	return &AuthHandler{
		loginUseCase: loginUsecase,
	}
}

func (h *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	accessToken, refreshToken, err := h.loginUseCase.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials: %v", err)
	}

	return &authpb.LoginResponse{
		TokenType:    "Bearer",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
