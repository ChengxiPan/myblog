package util

import "regexp"

// check if email is in correct form
func IsValidEmail(email string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
