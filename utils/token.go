package utils

import (
	"math/rand"

	"github.com/google/uuid"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,._-%&()")

func GenerateToken(length int) string {
	token := make([]byte, length)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}

func NewID() string {
	return uuid.New().String()
}
