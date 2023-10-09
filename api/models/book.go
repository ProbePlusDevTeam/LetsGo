// models/book.go
package models

// Book represents a book data structure
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}
