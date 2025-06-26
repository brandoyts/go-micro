package main

import (
	"fmt"
	"log"
	"microservices/api-gateway/internal/infrastructure/authClient"
	"microservices/api-gateway/internal/infrastructure/bookingClient"
	"microservices/api-gateway/internal/infrastructure/userClient"
	"microservices/api-gateway/internal/interface/handlers"
	"microservices/api-gateway/internal/interface/routes"
	"net/http"
	"os"
)

func main() {
	httpPort := ":8000"

	user := userClient.New(os.Getenv("USER_DSN"))
	auth := authClient.New(os.Getenv("AUTH_DSN"))
	booking := bookingClient.New(os.Getenv("BOOKING_DSN"))

	handler := handlers.NewHandler(user, auth, booking)

	router := routes.NewRouter(handler)

	server := http.Server{
		Addr:    httpPort,
		Handler: router,
	}

	fmt.Println("API Gateway is running on port", httpPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}

}
