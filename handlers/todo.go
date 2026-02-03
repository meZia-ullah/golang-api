package handlers

import (
	"encoding/json"
	"goapi/db"
	"goapi/models"
	"net/http"
	"strconv"
	"strings"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "user_id query parameter is required", 400)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user_id", 400)
		return
	}

	var todos []models.Todo
	if err := db.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}

	if todo.UserID == 0 || todo.Title == "" {
		http.Error(w, "user_id and title are required", 400)
		return
	}

	if err := db.DB.Create(&todo).Error; err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todo/")
	todoID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var input models.Todo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := db.DB.
		Model(&models.Todo{}).
		Where("id = ?", todoID).
		Updates(&input)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	input.ID = uint(todoID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(input)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todo/")
	todoID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid todo ID", 400)
		return
	}

	if err := db.DB.Delete(&models.Todo{}, todoID).Error; err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted successfully"})
}
