// models/book.go
package models

// Book represents a book data structure
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

type Book1 struct{
    ID string `json:"id"`
    Title string `json:"name"`
    Price float32 `json:"price"`
}

type Book2 struct{
    ID string `json:"id"`
    Author string `json:"author"`
    Published_Date string `json:"date"` 
}
