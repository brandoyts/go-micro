package main

import (
	"log"
	"microservices/user/database"
	"microservices/user/internal/grpc"
	"microservices/user/internal/repository"
	"microservices/user/internal/service"
	"os"
	"os/signal"
)

func init() {
	dsn := os.Getenv("USER_DB_DSN")
	db, migrateErr := database.Migrate(dsn)
	if migrateErr != nil {
		log.Fatalf("❌ Migration failed: %v", migrateErr)
	}

	seedErr := database.Seed(db)
	if seedErr != nil {
		log.Fatalf("❌ Seeding failed: %v", seedErr)
	}
}

func main() {

	userDb := database.New()

	bookingRepository := repository.NewUserRepository(userDb.Db)

	bookingService := service.NewUserService(bookingRepository)

	server := grpc.NewGRPCServer(bookingService)
	server.Start(":50051")

	// Graceful shutdown logic
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
}
