package routes

import (
	"go-shortener/handlers"
	"go-shortener/repositories"
	"go-shortener/services"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/signup", authHandler.SignUp)
		auth.POST("/signin", authHandler.SignIn)
	}
}
