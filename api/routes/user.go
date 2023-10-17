package routes

import (
    "github.com/gin-gonic/gin"

    "github.com/ProbePlusDevTeam/LetsGo/api/handlers"
)

// InitializeRoutes sets up all the routes for the application
func UserRoutes(user_route *gin.Engine) {
    user_route.POST("/create_user", handlers.CreateUSer)
	user_route.GET("/get_users",handlers.GetUsers)
    user_route.GET("/get_user/:id",handlers.GetUser)
	user_route.PUT("/update_user/:id",handlers.UpdateUser)
}