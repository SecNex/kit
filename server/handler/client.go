package handler

import (
	"encoding/json"
	"net/http"

	"github.com/secnex/kit/models"
)

type ClientGetResponse struct {
	Count int             `json:"count"`
	Apps  []models.Client `json:"clients"`
}

func (h *Handler) ClientGet(w http.ResponseWriter, r *http.Request) {
	var result []models.Client

	err := h.db.DB.Find(&result).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ClientGetResponse{
		Count: len(result),
		Apps:  result,
	})
}

func (h *Handler) ClientNew(w http.ResponseWriter, r *http.Request) {
	if !h.RequiredBodyFields(w, r, "name", "description", "slug", "application_id") {
		return
	}

	var client models.Client

	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		BadRequest(w, r, "Invalid request body")
		return
	}

	err = h.db.DB.Create(&client).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}
