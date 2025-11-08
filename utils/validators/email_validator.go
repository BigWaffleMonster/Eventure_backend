package validators

import "regexp"

const emailValidationRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

func IsValidEmail(email string) bool {
	return regexp.MustCompile(emailValidationRegex).MatchString(email)
}
