package handler

import (
	"encoding/json"
	"net/http"

	"github.com/secnex/kit/models"
)

type UserGetResponse struct {
	Count int           `json:"count"`
	Users []models.User `json:"users"`
}

type UserNewRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TenantID  string `json:"tenant_id"`
	Password  string `json:"password"`
}

func (h *Handler) UserGet(w http.ResponseWriter, r *http.Request) {
	var result []models.User

	err := h.db.DB.Find(&result).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserGetResponse{
		Count: len(result),
		Users: result,
	})
}

func (h *Handler) UserNew(w http.ResponseWriter, r *http.Request) {
	var request UserNewRequest

	if !h.RequiredBodyFields(w, r, "username", "email", "first_name", "last_name", "password", "tenant_id") {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		BadRequest(w, r, "Invalid request body")
		return
	}
	// Check if tenant is default ("default") then set tenant ID to 00000000-0000-0000-0000-000000000000
	if request.TenantID == "default" {
		request.TenantID = "00000000-0000-0000-0000-000000000000"
	}

	// Check if tenant exists
	var tenant models.Tenant
	err = h.db.DB.Where("id = ?", request.TenantID).First(&tenant).Error
	if err != nil {
		BadRequest(w, r, "Tenant not found!")
		return
	}

	// Check if username is already taken
	var user models.User
	err = h.db.DB.Where("username = ?", request.Username).First(&user).Error
	if err != nil {
		BadRequest(w, r, "Username already taken!")
		return
	}

	// Check if email is already taken
	err = h.db.DB.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		BadRequest(w, r, "Email already taken!")
		return
	}

	// Create user
	user = models.User{
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		TenantID:  request.TenantID,
		Password:  request.Password,
	}

	err = h.db.DB.Create(&user).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully! Please check your email for verification."})
}
