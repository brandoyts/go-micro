package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"microservices/pkg/proto-gen/userpb"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service userpb.UserServiceClient
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userUuid := chi.URLParam(r, "uuid")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	req := &userpb.GetUserRequest{Uuid: userUuid}
	res, err := h.Service.GetUser(ctx, req)
	if err != nil {
		http.Error(w, fmt.Sprintf("gRPC error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload userpb.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	res, err := h.Service.CreateUser(ctx, &payload)
	if err != nil {
		http.Error(w, "unable to create a user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
