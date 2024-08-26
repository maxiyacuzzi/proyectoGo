package routes

import (
    "myapp/controllers"
    "myapp/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    userRoutes := r.Group("/users")
    {
        userRoutes.GET("/", middleware.JWTMiddleware(), controllers.GetUsers)
        userRoutes.POST("/", middleware.JWTMiddleware(), controllers.CreateUser)
    }
}
