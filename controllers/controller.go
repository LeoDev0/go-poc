package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/go-poc/database"
	"example.com/go-poc/errors"
	"example.com/go-poc/models"
	"github.com/go-chi/chi"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
 	var books []models.Book

	database.DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book models.Book

	err := database.DB.First(&book, id)

	if err == nil {
		w.WriteHeader(errors.NotFound("").Status)
		json.NewEncoder(w).Encode(errors.NotFound(""))
		return
	}

	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)
	database.DB.Create(&book)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book models.Book

	w.WriteHeader(http.StatusNoContent)
	database.DB.Delete(&book, id)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book models.Book

	database.DB.First(&book, id)
	json.NewDecoder(r.Body).Decode(&book)

	database.DB.Save(&book)
	json.NewEncoder(w).Encode(book)
}