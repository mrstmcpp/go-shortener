package routes

import (
	"go-shortener/handlers"
	"go-shortener/repositories"
	"go-shortener/services"

	"github.com/gin-gonic/gin"
)

func RegisterURLRoutes(router *gin.Engine) {
	urlRepo := repositories.NewUrlRepository()
	urlService := services.NewURLService(urlRepo)
	urlHandler := handlers.NewURLHandler(urlService)

	urls := router.Group("/api/v1/urls")
	// urls.Use(middleware.)

	urls.POST("", urlHandler.CreateURL)
}
