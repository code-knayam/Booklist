package controllers

import (
	"Booklist/models"
	bookrepository "Booklist/repository/book"
	"Booklist/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Controller struct
type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks Handler func
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Get books is called")

		var book models.Book
		var error models.Error

		books = []models.Book{}

		bookRepo := bookrepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Server error"
			w.Header().Set("Content-Type", "application/json")
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

// GetBook Handler func
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Get book is called")

		var book models.Book
		var error models.Error

		params := mux.Vars(r)

		bookRepo := bookrepository.BookRepository{}
		id, _ := strconv.Atoi(params["id"])

		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "No rows found"
				w.Header().Set("Content-Type", "application/json")
				utils.SendError(w, http.StatusNotFound, error)
				return
			}
			error.Message = "Internal Server error"
			w.Header().Set("Content-Type", "application/json")
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

// AddBook Handler func
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Add book is called")

		var book models.Book
		var bookID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All fields are required"
			w.Header().Set("Content-Type", "application/json")
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookrepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Internal server error"
			w.Header().Set("Content-Type", "application/json")
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, bookID)
	}
}

// UpdateBook Handler func
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Update book is called")

		var book models.Book
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All fields are required"
			w.Header().Set("Content-Type", "application/json")
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookrepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Internal server error"
			w.Header().Set("Content-Type", "application/json")
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		if rowsUpdated == 0 {
			error.Message = "Book not found"
			w.Header().Set("Content-Type", "application/json")
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

// RemoveBook Handler Func
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Remove book is called")

		var error models.Error
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		bookRepo := bookrepository.BookRepository{}
		rowsDeleted, err := bookRepo.RemoveBook(db, id)

		if err != nil {
			error.Message = "Internal Server Error"
			w.Header().Set("Content-Type", "application/json")
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsDeleted == 0 {
			error.Message = "Record not found"
			w.Header().Set("Content-Type", "application/json")
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}
