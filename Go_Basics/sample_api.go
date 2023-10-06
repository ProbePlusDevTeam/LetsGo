//  qroute.GET("/hello", func(c *gin.Context) {
//         c.JSON(http.StatusOK, gin.H{
//             "message": "Hello, World!",
//         })
//     })

package main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

/*Function to get all the books details */
func getbooks(c *gin.Context){  //'c' represents context of incoming HTTP request containing information about request(headers,query parameters & response writer)
	
    c.JSON(http.StatusOK, books) //send an HTTP response in JSON format the HTTP status code & data to be serialized and sent as JSON
}

/*Function to get a book detail*/
func getbook(c *gin.Context){
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
func postbook(detail *gin.Context){
	var newbook Book
	if error := detail.BindJSON(&newbook); error!=nil{     //If the JSON data in the request body is successfully parsed & matches the structure of newbook variable, this function call will return nil (indicating no error).
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	}
	books= append(books,newbook)
	detail.JSON(http.StatusCreated,newbook)
}


func updatebook(detail *gin.Context) {
	id := detail.Param("id")
	var updatebook Book
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
func delbook(detail *gin.Context) {
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

func main() {
	r := gin.New()      //creates a new Gin router instance and assigns it to the variable r

	r.GET("/books", getbooks) 
	r.GET("/book/:id",getbook)
	r.POST("/create_book",postbook)
	r.PUT("/update_book/:id",updatebook)
	r.DELETE("/delete_book/:id",delbook)
	r.Run(":8080")
}