package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateID(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}
