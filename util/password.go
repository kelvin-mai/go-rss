package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatalf("Error generating password: %v\n", err)
	}
	return string(bytes)
}
