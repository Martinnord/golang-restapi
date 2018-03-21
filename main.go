package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book structure
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author structure
type Author struct {
	Firstname string `json:"fistname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init router
	r := mux.NewRouter()

	//mock data
	books = append(books, Book{ID: "1", Isbn: "4312", Title: "Hejsan", Author: &Author{Firstname: "Martin", Lastname: "Nordström"}})

	// Route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
