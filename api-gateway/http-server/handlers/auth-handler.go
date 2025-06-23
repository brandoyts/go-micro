package handlers

import (
	"context"
	"encoding/json"
	"microservices/pkg/proto-gen/authpb"
	"net/http"
	"time"
)

type AuthHandler struct {
	Service authpb.AuthServiceClient
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var raw struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	payload := authpb.LoginRequest{
		Username: raw.Username,
		Password: raw.Password,
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	res, err := a.Service.Login(ctx, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
