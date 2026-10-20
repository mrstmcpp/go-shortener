package handlers

import (
	"go-shortener/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
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
