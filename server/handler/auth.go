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
}

func (h *Handler) AuthLogout(w http.ResponseWriter, r *http.Request) {

}
