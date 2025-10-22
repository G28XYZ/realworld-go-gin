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
 }
}

func (h *UserHandler) Register(c *gin.Context) {
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

	user, err := h.usecase.Register(req.User.Email, req.User.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}