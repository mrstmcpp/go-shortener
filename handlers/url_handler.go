package handlers

import (
	"errors"
	"fmt"
	"go-shortener/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type URLHandler struct {
	urlService *services.URLService
}

func NewURLHandler(urlService *services.URLService) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}

type CreateNewURLRequest struct {
	URL string `json:"url" binding:"required"`
}

func (h *URLHandler) CreateURL(c *gin.Context) {
	var req CreateNewURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "A valid URL is required.",
			},
		)
		return
	}

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id, ok := userID.(bson.ObjectID)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authenticated user",
		})
		return
	}

	record, err := h.urlService.CreateNewUrl(req.URL, id)

	if err != nil {
		if err.Error() == "invalid URL" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create short URL",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Short URL created",
		"data":    record,
	})
}

func (h *URLHandler) RedirectHandler(c *gin.Context) {
	code := c.Param("shortCode")

	fmt.Println("code is here")

	record, err := h.urlService.RedirectURL(code)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Short URL not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not redirect url",
		})
		return
	}
	c.Redirect(http.StatusFound, record.OriginalUrl)
}
