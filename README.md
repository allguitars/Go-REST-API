# Golang: REST API
Implementation of REST APIs in Golang -- a book management system that has five route handlers:
- Get all books (GET)
- Get a single book with ID (GET)
- Create a new book (POST)
- Update the book with ID (PUT)
- Delete a book with ID (DELETE)

## Endpoints
```golang
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE") 
```

## Mock Data
```golang
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
```
 

## Tutorial
Check out [Traversy Media](https://www.youtube.com/user/TechGuyWeb)
