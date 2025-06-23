package grpcClient

import (
	"microservices/pkg/proto-gen/bookingpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitBookingService(addr string) (bookingpb.BookingServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := bookingpb.NewBookingServiceClient(conn)

	return client, nil
}
