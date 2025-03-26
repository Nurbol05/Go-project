package handlers

import (
	"RestAPI/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

var books []models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Книга не найдена", http.StatusNotFound)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	if strings.TrimSpace(book.Title) == "" {
		http.Error(w, "Название книги не может быть пустым", http.StatusBadRequest)
		return
	}
	if book.AuthorID <= 0 {
		http.Error(w, "Некорректный ID автора", http.StatusBadRequest)
		return
	}
	if book.CategoryID <= 0 {
		http.Error(w, "Некорректный ID категории", http.StatusBadRequest)
		return
	}
	if book.Price <= 0 {
		http.Error(w, "Цена должна быть больше 0", http.StatusBadRequest)
		return
	}

	book.ID = len(books) + 1
	books = append(books, book)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, book := range books {
		if book.ID == id {
			var updatedBook models.Book
			err := json.NewDecoder(r.Body).Decode(&updatedBook)

			if err != nil {
				http.Error(w, `{"error": "Ошибка в JSON"}`, http.StatusBadRequest)
				return
			}

			if updatedBook.Title == "" {
				http.Error(w, `{"error": "Название книги не может быть пустым"}`, http.StatusBadRequest)
				return
			}

			if updatedBook.AuthorID <= 0 {
				http.Error(w, `{"error": "ID автора должен быть больше 0"}`, http.StatusBadRequest)
				return
			}
			if updatedBook.CategoryID <= 0 {
				http.Error(w, `{"error": "ID категории должен быть больше 0"}`, http.StatusBadRequest)
				return
			}

			if updatedBook.Price <= 0 {
				http.Error(w, `{"error": "Цена книги должна быть больше 0"}`, http.StatusBadRequest)
				return
			}

			updatedBook.ID = id
			books[i] = updatedBook
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Книга не найдена"})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Книга не найдена", http.StatusNotFound)
}
