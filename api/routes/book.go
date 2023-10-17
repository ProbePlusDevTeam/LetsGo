// routes/routes.go
package routes

import (
    "github.com/gin-gonic/gin"

    "github.com/ProbePlusDevTeam/LetsGo/api/handlers"
)

// InitializeRoutes sets up all the routes for the application
func BookRoutes(book_route *gin.Engine) {
    book_route.GET("/books", handlers.GetBooks)
    book_route.GET("/book/:id", handlers.GetBook)
    book_route.POST("/create_book", handlers.PostBook)
    book_route.PUT("/update_book/:id", handlers.UpdateBook)
    book_route.DELETE("/delete_book/:id", handlers.DelBook)

    book_route.POST("/create_book1",handlers.PostBook1)
    book_route.POST("/create_book2",handlers.PostBook2)
    book_route.GET("/get_book_1",handlers.GetDetails)
    book_route.GET("/create_new_cn",handlers.CreateNewCollection)
}
