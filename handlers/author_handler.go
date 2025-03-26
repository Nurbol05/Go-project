package handlers

import (
	"RestAPI/models"
	"encoding/json"
	"net/http"
	"strings"
)

var authors []models.Author

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)

	if strings.TrimSpace(author.Name) == "" {
		http.Error(w, "Имя автора не может быть пустым", http.StatusBadRequest)
		return
	}

	author.ID = len(authors) + 1
	authors = append(authors, author)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}
