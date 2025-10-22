package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Создание нового пользователя (валидация + хэш)
func NewUser(username, email, password string) (*User, error) {

	if username == "" || email == "" || password == "" {
		return nil, errors.New("all fields are required")
	}

	address, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: username,
		Email:    address,
		Password: string(hashed),
	}, nil
}
