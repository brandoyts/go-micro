package main

import (
	"context"
	"fmt"
	"log"
	grpcClient "microservices/api-gateway/grpc-client"
	httpServer "microservices/api-gateway/http-server"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	// Handle interrupt signal
	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	userDsn := os.Getenv("USER_DSN")
	if userDsn == "" {
		log.Fatal("USER_DSN environment variable not set")
	}

	userService, initUserServiceErr := grpcClient.InitUserService(userDsn)
	if initUserServiceErr != nil {
		log.Fatal(initUserServiceErr)
	}

	bookingDsn := os.Getenv("BOOKING_DSN")
	if bookingDsn == "" {
		log.Fatal("BOOKING_DSN environment variable not set")
	}
	bookingService, initBookingServiceErr := grpcClient.InitBookingService(bookingDsn)
	if initBookingServiceErr != nil {
		log.Fatal(initBookingServiceErr)
	}

	authDsn := os.Getenv("AUTH_DSN")
	if authDsn == "" {
		log.Fatal("AUTH_DSN environment variable not set")
	}
	authService, initAuthServiceErr := grpcClient.InitAuthService(authDsn)
	if initAuthServiceErr != nil {
		log.Fatal(initAuthServiceErr)
	}

	// Create HTTP server
	srv := httpServer.New(httpServer.Config{
		Port:                 httpPort,
		ReadTimeout:          httpTimeoutDuration,
		WriteTimeout:         httpTimeoutDuration,
		IdleTimeout:          httpIdleTimeoutDuration,
		UserServiceClient:    userService,
		BookingServiceClient: bookingService,
		AuthServiceClient:    authService,
	})

	// Start HTTP server
	go func() {
		fmt.Println("API Gateway is running on port", httpPort)
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-shutdownCtx.Done()

	fmt.Println("Shutting down API Gateway...")
	if err := srv.Shutdown(httpTimeoutDuration); err != nil {
		log.Fatalf("Graceful shutdown failed: %v", err)
	}
	fmt.Println("API Gateway stopped.")
}
