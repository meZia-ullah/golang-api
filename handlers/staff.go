package handlers

import (
	"encoding/json"
	"goapi/db"
	"goapi/models"
	"net/http"
)

func GetStaff(w http.ResponseWriter, r *http.Request) {
	var staff []models.Staff
	if err := db.DB.Find(&staff).Error; err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(staff)
}

func CreateStaff(w http.ResponseWriter, r *http.Request) {
	var staff models.Staff
	json.NewDecoder(r.Body).Decode(&staff)

	if err := db.DB.Create(&staff).Error; err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(staff)
}
