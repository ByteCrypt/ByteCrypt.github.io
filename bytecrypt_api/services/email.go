package services

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

const EMAIL_EXPRESSION = `^[^\s@]+@[^\s@]+\.[^\s@]+$`

func (provider *Provider) ValidateEmail(email string) error {
	emailExp, err := regexp.Compile(EMAIL_EXPRESSION)
	if err != nil {
		return err
	}

	if !emailExp.MatchString(email) {
		return fmt.Errorf("email is invalid: %s", email)
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return fmt.Errorf("email is not correctly formatted: %s", email)
	}

	domain := parts[1]
	mx, err := net.LookupMX(domain)
	if err != nil || len(mx) == 0 {
		return fmt.Errorf("email domain is invalid %s", email)
	}

	return nil
}
