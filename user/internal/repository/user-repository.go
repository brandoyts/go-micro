package repository

import (
	"context"
	"microservices/pkg/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	FindByUuid(ctx context.Context, uuid string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	Find(ctx context.Context, user *entities.User) (*entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *userRepository) FindByUuid(ctx context.Context, uuid string) (*entities.User, error) {
	var user entities.User

	result := r.db.WithContext(ctx).First(&user, "uuid = ?", uuid)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	result := r.db.WithContext(ctx).
		Where("uuid = ?", user.Uuid).
		Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) Find(ctx context.Context, user *entities.User) (*entities.User, error) {
	result := r.db.WithContext(ctx).Where(user).First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
