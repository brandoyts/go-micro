package routes

import (
	"microservices/api-gateway/http-server/handlers"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(h *handlers.AuthHandler) chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Post("/login", h.Login)
	})

	return router
}
