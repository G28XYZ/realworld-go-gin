package auth

import (
	"fmt"
	"net/http"
	"time"

	"realworld-go-gin/src/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var cfg  = config.GetConfig()

var jwtSecret = []byte(cfg.Jwt.Phrase)

type UserAuth struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
}

type AuthRequest struct {
	User UserAuth `json:"user" binding:"required"`
}

type LoginResponse struct {
	User UserAuth `json:"user"`
}

type RegistrationResponse struct {
	User UserAuth `json:"user"`
}

func bindJSON(c *gin.Context, req interface{}) error {
    if err := c.ShouldBindJSON(req); err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid request"})
        return err
    }
    return nil
}


func createJwt(user UserAuth) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": user.Email,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    })

    // Не забываем подписать токен!
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", fmt.Errorf("failed to sign token: %w", err)
    }

    return tokenString, nil
}


func Login(c *gin.Context) {
	var req AuthRequest

	if err := bindJSON(c, &req); err != nil {
        return
    }

	email    := req.User.Email
	password := req.User.Password

	// TODO: Проверка пользователя в БД
	if email != "jake@jake.jake" || password != "jakejake" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}



	tokenString, err := createJwt(req.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	var resp LoginResponse
	resp.User.Email = email
	resp.User.Token = tokenString

	c.JSON(http.StatusOK, resp)
}


func Registration(c *gin.Context) {
	var req AuthRequest

	if err := bindJSON(c, &req); err != nil {
        return
    }

	email    := req.User.Email
	password := req.User.Password
	username := req.User.Username

	if email == "" || password == "" || username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	var resp struct {
		Success bool `json:"success"`
		Data string `json:"data"`
	}

	resp.Success = true
	resp.Data = "hello " + username

	c.JSON(http.StatusOK, resp)
}