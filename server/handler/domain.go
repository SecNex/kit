package handler

import (
	"encoding/json"
	"net/http"

	"github.com/secnex/kit/models"
)

type DomainGetResponse struct {
	Count   int             `json:"count"`
	Domains []models.Domain `json:"domains"`
}

type DomainNewRequest struct {
	Name           string `json:"name"`
	OrganizationID string `json:"organization_id"`
}

func (h *Handler) DomainGet(w http.ResponseWriter, r *http.Request) {
	var result []models.Domain

	err := h.db.DB.Find(&result).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(DomainGetResponse{
		Count:   len(result),
		Domains: result,
	})
}

func (h *Handler) DomainNew(w http.ResponseWriter, r *http.Request) {
	var result models.Domain

	if !h.RequiredBodyFields(w, r, "name", "organization_id") {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		BadRequest(w, r, "Invalid request body")
		return
	}

	err = h.db.DB.Create(&result).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
