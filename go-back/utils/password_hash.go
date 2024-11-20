package utils

import "crypto/sha512"

func HashPassword(password string, secretKey string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hasher.Write([]byte(secretKey))
	return string(hasher.Sum(nil))
}
