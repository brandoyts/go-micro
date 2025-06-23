package handlers

import (
	"context"
	"encoding/json"
	"microservices/pkg/proto-gen/bookingpb"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingHandler struct {
	Service bookingpb.BookingServiceClient
}

func (b *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	bookingUuid := chi.URLParam(r, "bookingUuid")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	res, err := b.Service.GetBooking(ctx, &bookingpb.GetBookingRequest{Uuid: bookingUuid})
	if err != nil {
		http.Error(w, "unable to get booking", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (b *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var raw struct {
		UserUuid string    `json:"user_uuid"`
		Schedule time.Time `json:"schedule"`
	}

	err := json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	payload := bookingpb.CreateBookingRequest{
		UserUuid: raw.UserUuid,
		Schedule: timestamppb.New(raw.Schedule),
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	res, err := b.Service.CreateBooking(ctx, &payload)
	if err != nil {
		http.Error(w, "unable to create a booking", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (b *BookingHandler) GetBookingsByUserUuid(w http.ResponseWriter, r *http.Request) {
	userUuid := chi.URLParam(r, "userUuid")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	res, err := b.Service.GetBookingsByUserUuid(ctx, &bookingpb.GetBookingsByUserUuidRequest{
		UserUuid: userUuid,
	})
	if err != nil {
		http.Error(w, "unable to get bookings", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (b *BookingHandler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	var raw struct {
		Uuid     string    `json:"uuid"`
		Schedule time.Time `json:"schedule"`
	}

	err := json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	payload := bookingpb.UpdateBookingRequest{
		Uuid:     raw.Uuid,
		Schedule: timestamppb.New(raw.Schedule),
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	res, err := b.Service.UpdateBooking(ctx, &payload)
	if err != nil {
		http.Error(w, "unable to update booking", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
