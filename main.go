package main

import (
	"go-shortener/config"
	"go-shortener/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	routes.RegisterAuthRoutes(router)
	routes.RegisterURLRoutes(router)

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}
