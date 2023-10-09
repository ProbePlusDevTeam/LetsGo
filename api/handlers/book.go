package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/ProbePlusDevTeam/LetsGo/config"
    "github.com/ProbePlusDevTeam/LetsGo/api/models"
)

var books = []models.Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}


/*Function to get all the books details */
func GetBooks(c *gin.Context){  //'c' represents context of incoming HTTP request containing information about request(headers,query parameters & response writer)
	
    c.JSON(http.StatusOK, books) //send an HTTP response in JSON format the HTTP status code & data to be serialized and sent as JSON
}

/*Function to get a book detail*/
func GetBook(c *gin.Context){
	id:= c.Param("id")
	for _,a := range books{
		if a.ID == id{
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound,gin.H{"msg":"Book Not Found"})
}

/*function to create a Book detail */
func PostBook(detail *gin.Context){
	var newbook models.Book
	if error := detail.BindJSON(&newbook); error!=nil{     //If the JSON data in the request body is successfully parsed & matches the structure of newbook variable, this function call will return nil (indicating no error).
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	}
	// import pdb;pdb.set_trace()

	if err := database.InsertBook(&newbook); err != nil {
        detail.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert into the database"})
        return
    }
	detail.JSON(http.StatusCreated,newbook)
}


func UpdateBook(detail *gin.Context) {
	id := detail.Param("id")
	var updatebook models.Book
	if error := detail.BindJSON(&updatebook); error!=nil{     
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	for i, data := range books {
        if data.ID == id {
			books[i] =updatebook
			detail.JSON(http.StatusOK,books[i])
			return
		}
	}
	detail.JSON(http.StatusNotFound,gin.H{"msg":"Book Not Found"})
}

/*Function to delete a book detail */
func DelBook(detail *gin.Context) {
    id := detail.Param("id")
    for i, data := range books {
        if data.ID == id {
            // Remove the book from the slice using slice reassignment
            books = append(books[:i], books[i+1:]...)
            detail.JSON(http.StatusOK,gin.H{"message": "Details successfully deleted"})
			return
        }
    }
	detail.JSON(http.StatusNotFound,gin.H{"msg":"Book Not Found"})
}