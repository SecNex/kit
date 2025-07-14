package handler

import (
	"encoding/json"
	"net/http"

	"github.com/secnex/kit/models"
	"github.com/secnex/kit/utils"
)

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}

type AuthRegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TenantID  string `json:"tenant_id"`
}

func (h *Handler) AuthLogin(w http.ResponseWriter, r *http.Request) {
	var request AuthLoginRequest

	json.NewDecoder(r.Body).Decode(&request)

	if request.Email == "" || request.Password == "" {
		BadRequest(w, r, "Email and password are required!")
		return
	}

	var user models.User

	err := h.db.DB.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		BadRequest(w, r, "Invalid email or password!")
		return
	}

	match, err := utils.Verify(request.Password, user.Password)
	if err != nil {
		Unauthorized(w, r)
		return
	}

	if !match {
		Unauthorized(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

func (h *Handler) AuthRegister(w http.ResponseWriter, r *http.Request) {
	var request AuthRegisterRequest

	json.NewDecoder(r.Body).Decode(&request)

	if request.Username == "" || request.Email == "" || request.Password == "" || request.FirstName == "" || request.LastName == "" || request.TenantID == "" {
		BadRequest(w, r, "Username, email, password, first name, last name, and tenant ID are required!")
		return
	}

	// Check if tenant exists
	var tenant models.Tenant
	err := h.db.DB.Where("id = ?", request.TenantID).First(&tenant).Error
	if err != nil {
		BadRequest(w, r, "Invalid tenant ID!")
		return
	}

	// Check if email is already in use
	var user models.User
	err = h.db.DB.Where("email = ? AND tenant_id = ?", request.Email, request.TenantID).First(&user).Error
	if err == nil {
		BadRequest(w, r, "User already exists!")
		return
	}

	// Check if username is already in use
	err = h.db.DB.Where("username = ? AND tenant_id = ?", request.Username, request.TenantID).First(&user).Error
	if err == nil {
		BadRequest(w, r, "Username already exists!")
		return
	}

	// Create user
	user = models.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		TenantID:  request.TenantID,
	}

	err = h.db.DB.Create(&user).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func (h *Handler) AuthLogout(w http.ResponseWriter, r *http.Request) {

}
