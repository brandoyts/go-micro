package application

import (
	"context"
	"errors"
	"microservices/auth/internal/infrastructure/jwt"
	"microservices/auth/internal/infrastructure/userClient"

	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	userClient userClient.UserClient
	jwtService jwt.JwtService
}

func NewLoginUsecase(userClient userClient.UserClient, jwtService jwt.JwtService) *LoginUsecase {
	return &LoginUsecase{
		userClient: userClient,
		jwtService: jwtService,
	}
}

func (uc *LoginUsecase) Login(ctx context.Context, username, password string) (accessToken, refreshToken string, err error) {
	// Step 1: Fetch user info from user service via gRPC
	user, err := uc.userClient.GetUserByUsername(ctx, username)
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	// Step 2: Compare password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return "", "", errors.New("invalid username or password")
	}

	// Step 3: Generate JWT tokens
	accessToken, err = uc.jwtService.GenerateAccessToken(user.Uuid)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = uc.jwtService.GenerateRefreshToken(user.Uuid)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
