package handlers

import (
	"encoding/json"
	"goapi/db"
	"goapi/models"
	"strconv"
	"strings"

	"net/http"
)

func GetEducations(w http.ResponseWriter, r *http.Request) {
	var educations []models.Education
	if err := db.DB.Find(&educations).Error; err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(educations)
}

func CreateEducation(w http.ResponseWriter, r *http.Request) {
	var edu models.Education
	if err := json.NewDecoder(r.Body).Decode(&edu); err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}
	if edu.StudentName == "" || edu.Course == "" || edu.Grade == "" {
		http.Error(w, "please fill in the fields", http.StatusBadRequest)
		return
	}
	if err := db.DB.Create(&edu).Error; err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(edu)
}
func UpdateEducation(w http.ResponseWriter, r *http.Request) {
	IdStr := strings.TrimPrefix(r.URL.Path, "/Education/")
	id, err := strconv.ParseUint(IdStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var user models.Education
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid body formate", http.StatusBadRequest)
		return
	}
	if err := db.DB.Model(&models.Education{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	user.ID = uint(id)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}
