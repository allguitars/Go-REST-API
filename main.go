package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/* Models
----------------------------------------------------------------------- */

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"` // Fetch the id property in JSON
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"` // Author has its own struct
}

// Author struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice of Book struct
var books []Book

/* Route Handlers
----------------------------------------------------------------------- */
// Get all Books
// Every route handler shoule take in an http.ResponseWriter and the pointer
// of http.Request.
func getBooks(w http.ResponseWriter, r *http.Request) {

	// The content will be sent as plain text if you don't set this content type
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	// Decode the request body and put the content to book
	// _ is a convention for doing nothing with the return value
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	// Append the new book to the collection
	books = append(books, book)

	// Return the new information to the browser
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {

		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) // Delete the oroginal one

			// Create a new one with the same ID
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init the router
	r := mux.NewRouter()

	// Mock data - @todo: implement DB
	books = append(books, Book{
		ID:     "1",
		Isbn:   "33456",
		Title:  "Book One",
		Author: &Author{Firstname: "Jane", Lastname: "Doe"}})
	books = append(books, Book{
		ID:     "2",
		Isbn:   "54678",
		Title:  "Book Two",
		Author: &Author{Firstname: "David", Lastname: "Tao"}})

	// Creare Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run server
	fmt.Println("Server starting on port:8888...")
	log.Fatal(http.ListenAndServe(":8888", r))

}
