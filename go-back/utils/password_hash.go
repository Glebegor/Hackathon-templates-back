package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashPassword(password string, secretKey string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hasher.Write([]byte(secretKey))
	// Encode the raw hash bytes into a hexadecimal string
	return hex.EncodeToString(hasher.Sum(nil))
}
