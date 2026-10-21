package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const secretKey = "aslfdkaskflajkflajlfkdjsldkfj"

func GetJWTSecret() string {
	return secretKey
}

func CreateToken(userId bson.ObjectID) (string, error) {
	secretKey := secretKey
	fmt.Println("secret key " + secretKey)
	if secretKey == "" {
		return "", errors.New("JWT_SECRET is not configured")
	}

	claims := jwt.RegisteredClaims{
		Subject:   userId.Hex(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
