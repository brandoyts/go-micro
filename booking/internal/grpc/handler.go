package grpc

import (
	"context"
	"log"
	"microservices/pkg/entities"
	"microservices/pkg/proto-gen/bookingpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetBooking(ctx context.Context, in *bookingpb.GetBookingRequest) (*bookingpb.GetBookingResponse, error) {
	result, err := s.bookingService.FindByUuid(ctx, in.Uuid)

	if err != nil {
		log.Fatal(err)
	}

	return &bookingpb.GetBookingResponse{
		Booking: &bookingpb.Booking{
			Uuid:      result.Uuid,
			UserUuid:  result.UserUuid,
			Schedule:  timestamppb.New(result.Schedule),
			CreatedAt: timestamppb.New(result.CreatedAt),
			UpdatedAt: timestamppb.New(result.UpdatedAt),
		},
	}, nil
}

func (s *Server) GetBookingsByUserUuid(ctx context.Context, in *bookingpb.GetBookingsByUserUuidRequest) (*bookingpb.GetBookingsByUserUuidResponse, error) {
	result, err := s.bookingService.FindBookingByUserUuid(ctx, in.UserUuid)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get bookings: %v", err)
	}

	var bookings []*bookingpb.Booking

	for _, b := range result {
		bookings = append(bookings, &bookingpb.Booking{
			Uuid:      b.Uuid,
			UserUuid:  b.UserUuid,
			Schedule:  timestamppb.New(b.Schedule),
			CreatedAt: timestamppb.New(b.CreatedAt),
			UpdatedAt: timestamppb.New(b.UpdatedAt),
		})
	}

	return &bookingpb.GetBookingsByUserUuidResponse{
		Bookings: bookings,
	}, nil
}

func (s *Server) CreateBooking(ctx context.Context, in *bookingpb.CreateBookingRequest) (*bookingpb.CreateBookingResponse, error) {
	payload := entities.Booking{
		UserUuid: in.UserUuid,
		Schedule: in.Schedule.AsTime(),
	}

	result, err := s.bookingService.CreateBooking(ctx, &payload)

	if err != nil {
		log.Fatal(err)
	}

	return &bookingpb.CreateBookingResponse{
		Booking: &bookingpb.Booking{
			Uuid:      result.Uuid,
			UserUuid:  result.UserUuid,
			Schedule:  timestamppb.New(result.Schedule),
			CreatedAt: timestamppb.New(result.CreatedAt),
			UpdatedAt: timestamppb.New(result.UpdatedAt),
		},
	}, nil
}

func (s *Server) UpdateBooking(ctx context.Context, in *bookingpb.UpdateBookingRequest) (*bookingpb.UpdateBookingResponse, error) {
	payload := entities.Booking{
		Uuid:     in.Uuid,
		Schedule: in.Schedule.AsTime(),
	}

	result, err := s.bookingService.UpdateBooking(ctx, &payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update booking: %v", err)
	}

	return &bookingpb.UpdateBookingResponse{
		Booking: &bookingpb.Booking{
			Uuid:      result.Uuid,
			UserUuid:  result.UserUuid,
			Schedule:  timestamppb.New(result.Schedule),
			CreatedAt: timestamppb.New(result.CreatedAt),
			UpdatedAt: timestamppb.New(result.UpdatedAt),
		},
	}, nil
}
