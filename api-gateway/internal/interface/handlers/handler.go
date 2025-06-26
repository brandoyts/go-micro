package handlers

import (
	"context"
	"encoding/json"
	"microservices/api-gateway/internal/infrastructure/authClient"
	"microservices/api-gateway/internal/infrastructure/bookingClient"
	"microservices/api-gateway/internal/infrastructure/userClient"
	"microservices/pkg/proto-gen/authpb"
	"microservices/pkg/proto-gen/bookingpb"
	"microservices/pkg/proto-gen/userpb"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	User    userClient.Client
	Auth    authClient.Client
	Booking bookingClient.Client
}

func NewHandler(user userClient.Client, auth authClient.Client, booking bookingClient.Client) *Handler {
	return &Handler{
		User:    user,
		Auth:    auth,
		Booking: booking,
	}
}

func (h *Handler) jsonResponse(w http.ResponseWriter, data interface{}) {
	if data == nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	r.Context()
	w.Write([]byte("OK"))
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	payload := &userpb.GetUserRequest{
		Uuid: chi.URLParam(r, "uuid"),
	}

	response, err := h.User.GetUser(ctx, payload)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	h.jsonResponse(w, response)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var payload userpb.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.User.CreateUser(ctx, &payload)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	h.jsonResponse(w, response)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var payload authpb.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.Auth.Login(ctx, &payload)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	h.jsonResponse(w, response)
}

func (h *Handler) GetBooking(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	uuid := chi.URLParam(r, "uuid")

	payload := bookingpb.GetBookingRequest{
		Uuid: uuid,
	}

	response, err := h.Booking.GetBooking(ctx, &payload)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	h.jsonResponse(w, response)
}

func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var raw struct {
		UserUuid string    `json:"user_uuid"`
		Schedule time.Time `json:"schedule"`
	}

	decodeErr := json.NewDecoder(r.Body).Decode(&raw)
	if decodeErr != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	payload := bookingpb.CreateBookingRequest{
		UserUuid: raw.UserUuid,
		Schedule: timestamppb.New(raw.Schedule),
	}

	response, err := h.Booking.CreateBooking(ctx, &payload)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	h.jsonResponse(w, response)
}

func (h *Handler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var raw struct {
		Uuid     string    `json:"uuid"`
		Schedule time.Time `json:"schedule"`
	}

	decodeErr := json.NewDecoder(r.Body).Decode(&raw)
	if decodeErr != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	payload := bookingpb.UpdateBookingRequest{
		Uuid:     raw.Uuid,
		Schedule: timestamppb.New(raw.Schedule),
	}

	response, err := h.Booking.UpdateBooking(ctx, &payload)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	h.jsonResponse(w, response)
}

func (h *Handler) GetBookingByUserUuid(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	userUuid := chi.URLParam(r, "userUuid")

	payload := bookingpb.GetBookingsByUserUuidRequest{
		UserUuid: userUuid,
	}

	response, err := h.Booking.GetBookingByUserUuid(ctx, &payload)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	h.jsonResponse(w, response)
}
