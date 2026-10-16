package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateToken(userId bson.ObjectID) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")

	if secretKey == "" {
		return "", errors.New("JWT_SECRET is not configured")
	}

	claims := jwt.RegisteredClaims{
		Subject:   userId.Hex(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token.Claims)
	return token.SignedString([]byte(secretKey))
}
