package bookingClient

import (
	"context"
	"fmt"
	"log"
	"microservices/pkg/proto-gen/bookingpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	GetBooking(ctx context.Context, in *bookingpb.GetBookingRequest) (*bookingpb.GetBookingResponse, error)
	CreateBooking(ctx context.Context, in *bookingpb.CreateBookingRequest) (*bookingpb.CreateBookingResponse, error)
	GetBookingByUserUuid(ctx context.Context, in *bookingpb.GetBookingsByUserUuidRequest) (*bookingpb.GetBookingsByUserUuidResponse, error)
	UpdateBooking(ctx context.Context, in *bookingpb.UpdateBookingRequest) (*bookingpb.UpdateBookingResponse, error)
}

type client struct {
	client bookingpb.BookingServiceClient
}

func New(address string) Client {
	connection, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect to booking service at %s: %v", address, err)
	}

	return &client{
		client: bookingpb.NewBookingServiceClient(connection),
	}
}

func (bc *client) GetBooking(ctx context.Context, in *bookingpb.GetBookingRequest) (*bookingpb.GetBookingResponse, error) {
	response, err := bc.client.GetBooking(ctx, in)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func (bc *client) CreateBooking(ctx context.Context, in *bookingpb.CreateBookingRequest) (*bookingpb.CreateBookingResponse, error) {
	response, err := bc.client.CreateBooking(ctx, in)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func (bc *client) GetBookingByUserUuid(ctx context.Context, in *bookingpb.GetBookingsByUserUuidRequest) (*bookingpb.GetBookingsByUserUuidResponse, error) {
	response, err := bc.client.GetBookingsByUserUuid(ctx, in)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func (bc *client) UpdateBooking(ctx context.Context, in *bookingpb.UpdateBookingRequest) (*bookingpb.UpdateBookingResponse, error) {
	response, err := bc.client.UpdateBooking(ctx, in)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}
