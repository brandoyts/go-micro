package service

import (
	"context"
	"microservices/pkg/entities"
	"microservices/user/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	FindByUuid(ctx context.Context, uuid string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	Find(ctx context.Context, user *entities.User) (*entities.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, in *entities.User) (*entities.User, error) {
	return s.repo.Create(ctx, in)
}

func (s *userService) FindByUuid(ctx context.Context, in string) (*entities.User, error) {
	return s.repo.FindByUuid(ctx, in)
}

func (s *userService) Update(ctx context.Context, in *entities.User) (*entities.User, error) {
	return s.repo.Update(ctx, in)
}

func (s *userService) Find(ctx context.Context, in *entities.User) (*entities.User, error) {
	return s.repo.Find(ctx, in)
}
