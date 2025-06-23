package routes

import (
	"microservices/api-gateway/http-server/handlers"

	"github.com/go-chi/chi/v5"
)

func BookingRouter(h *handlers.BookingHandler) chi.Router {
	bookingRouter := chi.NewRouter()

	bookingRouter.Route("/", func(r chi.Router) {
		r.Get("/{bookingUuid}", h.GetBooking)
		r.Post("/", h.CreateBooking)
		r.Get("/user/{userUuid}", h.GetBookingsByUserUuid)
		r.Put("/", h.UpdateBooking)
	})

	return bookingRouter
}
