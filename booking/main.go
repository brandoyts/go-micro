package main

import (
	"log"
	"microservices/booking/database"
	"microservices/booking/internal/grpc"
	"microservices/booking/internal/repository"
	"microservices/booking/internal/service"
	"os"
	"os/signal"
)

func init() {
	dsn := os.Getenv("BOOKING_DB_DSN")
	_, err := database.Migrate(dsn)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
}

func main() {

	bookingDb := database.New()

	bookingRepository := repository.NewBookingRepository(bookingDb.Db)

	bookingService := service.NewBookingService(bookingRepository)

	server := grpc.NewGRPCServer(bookingService)
	server.Start(":50052")

	// Graceful shutdown logic
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
}
