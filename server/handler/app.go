package handler

import (
	"encoding/json"
	"net/http"

	"github.com/secnex/kit/models"
)

type AppGetResponse struct {
	Count int                  `json:"count"`
	Apps  []models.Application `json:"apps"`
}

type AppNewResponse struct {
	ID string `json:"id"`
}

func (h *Handler) AppGet(w http.ResponseWriter, r *http.Request) {
	var result []models.Application

	err := h.db.DB.Find(&result).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AppGetResponse{
		Count: len(result),
		Apps:  result,
	})
}

func (h *Handler) AppNew(w http.ResponseWriter, r *http.Request) {
	if !h.RequiredBodyFields(w, r, "name", "description", "slug", "tenant_id") {
		return
	}

	var app models.Application

	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		BadRequest(w, r, "Invalid request body")
		return
	}

	var createdApp models.Application

	err = h.db.DB.Create(&app).Scan(&createdApp).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AppNewResponse{
		ID: createdApp.ID,
	})
}
