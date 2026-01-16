package validation

import (
	"net/mail"
	"strings"
)

func ValidEmail(email string) bool {
	email = strings.TrimSpace(email)

	if email == "" {
		return false
	}

	_, err := mail.ParseAddress(email)
	return err == nil
}
