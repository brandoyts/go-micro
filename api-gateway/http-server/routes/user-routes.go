package routes

import (
	"microservices/api-gateway/http-server/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRouter(h *handlers.UserHandler) chi.Router {
	userRouter := chi.NewRouter()

	userRouter.Route("/", func(r chi.Router) {
		r.Get("/{uuid}", h.GetUser)
		r.Post("/create", h.CreateUser)
	})

	return userRouter
}
