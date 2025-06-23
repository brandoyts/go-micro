package grpc

import (
	"log"
	"microservices/booking/internal/service"
	"microservices/pkg/proto-gen/bookingpb"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	bookingpb.UnimplementedBookingServiceServer
	bookingService service.BookingService
}

func NewGRPCServer(svc service.BookingService) *Server {
	return &Server{
		bookingService: svc,
	}
}

func (s *Server) Start(port string) {

	go func() {
		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen on port %v: %v", port, err)
		}

		grpcServer := grpc.NewServer()
		bookingpb.RegisterBookingServiceServer(grpcServer, s)

		log.Printf("gRPC server listening at %v", listener.Addr())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
