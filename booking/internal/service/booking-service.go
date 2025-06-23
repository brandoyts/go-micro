package service

import (
	"context"
	"microservices/booking/internal/repository"
	"microservices/pkg/entities"
)

type BookingService interface {
	CreateBooking(ctx context.Context, booking *entities.Booking) (*entities.Booking, error)
	FindBookingByUserUuid(ctx context.Context, userUuid string) ([]entities.Booking, error)
	FindByUuid(ctx context.Context, uuid string) (*entities.Booking, error)
	UpdateBooking(ctx context.Context, booking *entities.Booking) (*entities.Booking, error)
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{
		repo: repo,
	}
}

func (s *bookingService) CreateBooking(ctx context.Context, booking *entities.Booking) (*entities.Booking, error) {
	return s.repo.Create(ctx, booking)
}

func (s *bookingService) FindBookingByUserUuid(ctx context.Context, userUuid string) ([]entities.Booking, error) {
	return s.repo.FindByUserUuid(ctx, userUuid)
}

func (s *bookingService) FindByUuid(ctx context.Context, uuid string) (*entities.Booking, error) {
	return s.repo.FindByUuid(ctx, uuid)
}

func (s *bookingService) UpdateBooking(ctx context.Context, booking *entities.Booking) (*entities.Booking, error) {
	return s.repo.Update(ctx, booking)
}
