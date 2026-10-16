package handlers

import (
	"errors"
	"go-shortener/services"
	"go-shortener/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Valid name, email, and password (minimum 5 characters) are required",
		})
		return
	}

	user, err := h.authService.SignUp(
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {
		if errors.Is(err, utils.ErrEmailExists) {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email already registered",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create account",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account created successfully",
		"user":    user,
	})
}

func (s *AuthHandler) SignIn(c *gin.Context) {
	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Valid email, and password are required",
		})
		return
	}

	user, err := s.authService.SignIn(req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Could not find user with provided email",
			"error2": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"token":   user,
	})
}
