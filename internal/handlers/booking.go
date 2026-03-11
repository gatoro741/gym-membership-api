package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) BookClass(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(int64)
	if !ok {
		writeError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req BookingRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	booking, err := h.service.BookClass(r.Context(), userId, req.ClassId)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func (h *Handler) GetMyBookings(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(int64)
	if !ok {
		writeError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	bookings, err := h.service.GetMyBookings(r.Context(), userId)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookings)

}

func (h *Handler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(int64)
	if !ok {
		writeError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	bookingId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid booking id")
		return
	}

	err = h.service.CancelBooking(r.Context(), bookingId, userId)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
