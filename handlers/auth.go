package handlers

import (
	"encoding/json"
	"goapi/db"
	"goapi/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}

	if user.Email == "" || user.Password == "" || user.Username == "" {
		http.Error(w, "Email, password and username are required", 400)
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "User already exists", 409)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := models.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Message:  "User registered successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginReq models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", 404)
		return
	}

	// Simple password check (in production, use bcrypt)
	if user.Password != loginReq.Password {
		http.Error(w, "Invalid password", 401)
		return
	}

	if !user.IsActive {
		http.Error(w, "User is inactive", 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := models.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Message:  "Login successful",
	}
	json.NewEncoder(w).Encode(response)
}
