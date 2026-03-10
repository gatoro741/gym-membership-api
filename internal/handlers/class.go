package handlers

import (
	"GymMembership-api/internal/models"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var req ClassRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	class := &models.Class{
		Title:       req.Title,
		TrainerName: req.TrainerName,
		Capacity:    req.Capacity,
		StartTime:   req.StartTime,
	}

	err = h.service.CreateClass(r.Context(), class)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := h.service.GetAllClasses(r.Context())
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(classes)
}
