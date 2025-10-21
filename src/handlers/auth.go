package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Секрет для подписи JWT (в реальном проекте брать из .env)
var jwtSecret = []byte("your_secret_key")

type LoginRequest struct {
	User struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	} `json:"user" binding:"required"`
}

type LoginResponse struct {
	User struct {
		Email string `json:"email"`
		Token string `json:"token"`
	} `json:"user"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid request"})
		return
	}

	email := req.User.Email
	password := req.User.Password

	// TODO: Проверка пользователя в БД
	if email != "jake@jake.jake" || password != "jakejake" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Создаем JWT токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	var resp LoginResponse
	resp.User.Email = email
	resp.User.Token = tokenString

	c.JSON(http.StatusOK, resp)
}
