package repository

import (
	"context"
	"microservices/pkg/entities"

	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(ctx context.Context, booking *entities.Booking) (*entities.Booking, error)
	FindByUserUuid(ctx context.Context, userUuid string) ([]entities.Booking, error)
	FindByUuid(ctx context.Context, uuid string) (*entities.Booking, error)
	Update(ctx context.Context, booking *entities.Booking) (*entities.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(ctx context.Context, booking *entities.Booking) (*entities.Booking, error) {
	err := r.db.WithContext(ctx).Create(booking).Error
	if err != nil {
		return nil, err
	}

	return booking, err
}

func (r *bookingRepository) FindByUserUuid(ctx context.Context, userUuid string) ([]entities.Booking, error) {
	var booking []entities.Booking

	result := r.db.WithContext(ctx).
		Where("user_uuid = ?", userUuid).
		Find(&booking)

	if result.Error != nil {
		return nil, result.Error
	}

	return booking, nil
}

func (r *bookingRepository) FindByUuid(ctx context.Context, uuid string) (*entities.Booking, error) {
	var booking entities.Booking

	result := r.db.WithContext(ctx).First(&booking, "uuid = ?", uuid)

	if result.Error != nil {
		return nil, result.Error
	}

	return &booking, nil
}

func (r *bookingRepository) Update(ctx context.Context, booking *entities.Booking) (*entities.Booking, error) {
	result := r.db.WithContext(ctx).
		Where("uuid = ?", booking.Uuid).
		Save(booking)
	if result.Error != nil {
		return nil, result.Error
	}

	return booking, nil
}
