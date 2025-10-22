package http

import (
	"net/http"
	"realworld-go-gin/internal/application"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase *application.UserUseCase
}

func NewUserHandler(r *gin.Engine, usecase *application.UserUseCase) {
	h := &UserHandler{usecase}
	api := r.Group("/api")
	{
		api.POST("/users", h.Register)
		api.POST("/users/login", h.Login)
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		User struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.usecase.Register(req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		User struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.usecase.Login(req.User.Email, req.User.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
