package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string, hashCost int) ([]byte, error) {
	plainPasswordAsByte := []byte(plainPassword)
	hash, err := bcrypt.GenerateFromPassword(plainPasswordAsByte, hashCost)

	return hash, err
}