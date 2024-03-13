package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), 8)

	return string(hashed)
}
