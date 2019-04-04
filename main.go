package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lib/pq"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	log.Println(pgURL)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get books is called")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get book is called")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book is called")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update book is called")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove book is called")
}
