package handler

import (
	"encoding/json"
	"net/http"

	"github.com/secnex/kit/models"
)

type TenantGetResponse struct {
	Count   int             `json:"count"`
	Tenants []models.Tenant `json:"tenants"`
}

type TenantNewRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organization_id"`
	DomainID       string `json:"domain_id"`
}

func (h *Handler) TenantGet(w http.ResponseWriter, r *http.Request) {
	var result []models.Tenant

	err := h.db.DB.Find(&result).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TenantGetResponse{
		Count:   len(result),
		Tenants: result,
	})
}

func (h *Handler) TenantNew(w http.ResponseWriter, r *http.Request) {
	var result models.Tenant

	if !h.RequiredBodyFields(w, r, "name", "description", "organization_id", "domain_id") {
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
