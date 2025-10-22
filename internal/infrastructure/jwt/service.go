package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	secret string
}

func NewService(secret string) *Service {
	return &Service{secret: secret}
}

func (s *Service) GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(72 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(s.secret))
}
