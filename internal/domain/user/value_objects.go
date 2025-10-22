package user

import (
	"errors"
	"regexp"
)

type Email struct {
	Value string
}

func NewEmail(email string) (Email, error) {
	re := regexp.MustCompile(`^[a-z0-9._%+]`)
	if !re.MatchString(email) {
		return Email{}, errors.New("invalid email format")
	}

	return Email{Value: email}, nil
}