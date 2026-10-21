package routes

import (
	"go-shortener/handlers"
	"go-shortener/middleware"
	"go-shortener/repositories"
	"go-shortener/services"

	"github.com/gin-gonic/gin"
)

func RegisterURLRoutes(router *gin.Engine) {
	urlRepo := repositories.NewUrlRepository()
	urlService := services.NewURLService(urlRepo)
	urlHandler := handlers.NewURLHandler(urlService)

	router.GET("/:shortCode", urlHandler.RedirectHandler)

	urls := router.Group("/api/v1/urls")
	urls.Use(middleware.AuthMiddlewareHelper())

	urls.POST("", urlHandler.CreateURL)
	urls.DELETE("/:id", urlHandler.DeleteURLByIDHandler)
	urls.GET("", urlHandler.FindByUserID)
}
