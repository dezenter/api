package utils

import (
	"log"
	"regexp"

	gonanoid "github.com/matoous/go-nanoid"
	"golang.org/x/crypto/bcrypt"
)

// GenerateID
func GenerateID(params ...int) string {
	var size = 20
	if len(params) == 1 {
		size = params[0]
	}
	id, err := gonanoid.Generate("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890", size)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

// HashPassword
func HashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(h), err
}

// CheckPasswordHash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidateEmail
func ValidateEmail(email string) bool {
	var r = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return r.MatchString(email)
}
