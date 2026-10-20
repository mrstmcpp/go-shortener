package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateShortCode() (string, error) {
	bytes := make([]byte, 6)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(bytes), nil
}
