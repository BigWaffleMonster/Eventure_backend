package helpers

import "regexp"

func IsValidEmail(email string) bool {
	// Регулярное выражение для проверки email
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
