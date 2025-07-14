package handler

import (
	"encoding/json"
	"net/http"

	"github.com/secnex/kit/models"
)

type OrganizationGetResponse struct {
	Count         int                   `json:"count"`
	Organizations []models.Organization `json:"organizations"`
}

type OrganizationNewRequest struct {
	Name string `json:"name"`
}

func (h *Handler) OrganizationGet(w http.ResponseWriter, r *http.Request) {
	var result []models.Organization

	err := h.db.DB.Find(&result).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(OrganizationGetResponse{
		Count:         len(result),
		Organizations: result,
	})
}

func (h *Handler) OrganizationNew(w http.ResponseWriter, r *http.Request) {
	var result models.Organization

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
