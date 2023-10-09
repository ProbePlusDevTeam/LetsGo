package main

import (
	"fmt"
	"context"
	"log"
    "github.com/gin-gonic/gin"
	"github.com/ProbePlusDevTeam/LetsGo/config"
    "github.com/ProbePlusDevTeam/LetsGo/api/routes"
)

func main() {
    // Create a Gin router instance
    book_route := gin.Default()
	client, err := database.ConnectMongoDB()
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

	defer client.Disconnect(context.Background())
	fmt.Printf("MongoDB Client: %+v\n", client)
    // Initialize routes from the routes file
    routes.BookRoutes(book_route)

    // Start the server
    book_route.Run(":8080")
}
