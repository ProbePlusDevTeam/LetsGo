package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ProbePlusDevTeam/LetsGo/api/routes"
	database "github.com/ProbePlusDevTeam/LetsGo/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Create a Gin router instance
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port := os.Getenv("PORT")

	client, err := database.ConnectMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	fmt.Printf("MongoDB Client: %+v\n", client)

	// Initialize routes from the routes file
	routes.BookRoutes(router)
	routes.UserRoutes(router)

	// Start the server
	router.Run(":" + port)
}
