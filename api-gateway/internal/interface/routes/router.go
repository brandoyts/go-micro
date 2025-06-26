package routes

import (
	"microservices/api-gateway/internal/interface/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	jwtSecret := "secret33210"
	tokenAuth := jwtauth.New("HS256", []byte(jwtSecret), nil)
	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {

		r.Group(func(protectedRoute chi.Router) {
			protectedRoute.Use(jwtauth.Verifier(tokenAuth))
			protectedRoute.Use(jwtauth.Authenticator(tokenAuth))

			protectedRoute.Route("/bookings", func(booking chi.Router) {
				booking.Get("/{uuid}", h.GetBooking)
				booking.Get("/user/{userUuid}", h.GetBookingByUserUuid)
				booking.Post("/create", h.CreateBooking)
				booking.Post("/update", h.UpdateBooking)
			})
		})

		r.Route("/users", func(u chi.Router) {
			u.With(jwtauth.Verifier(tokenAuth), jwtauth.Authenticator(tokenAuth)).Get("/{uuid}", h.GetUser)
			u.Post("/create", h.CreateUser)
		})

		r.Route("/auth", func(a chi.Router) {
			a.Post("/login", h.Login)
		})

		r.Get("/health", h.Health)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Custom 404: Route not found", http.StatusNotFound)
	})

	return router
}
