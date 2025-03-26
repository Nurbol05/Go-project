package handlers

import (
	"RestAPI/models"
	"encoding/json"
	"net/http"
)

var categories []models.Category

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)
	category.ID = len(categories) + 1
	categories = append(categories, category)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}
