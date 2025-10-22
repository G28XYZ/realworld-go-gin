package main

import (
	"fmt"
	"realworld-go-gin/src/config"
	"realworld-go-gin/src/handlers/auth"

	"github.com/gin-gonic/gin"
)


func main() {
	cfg := config.GetConfig()

	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/users/login", auth.Login)
		api.POST("/users", auth.Registration)
	}

	port := cfg.Server.Port
	fmt.Printf("\n🚀 Сервер запущен ➡️    http://localhost:%s    ⬅️\n", port)
	r.Run(":" + port)
}