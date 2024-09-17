package util

import (
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// check if email is in correct form
func IsValidEmail(email string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// hash and salt password: bycrypt
func HashAndSalt(pwd string) string {
	// Generate byscript hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error generating hash: %v", err)
	}
	return string(hashedPassword)
}

// compare password with hashed password
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
