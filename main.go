package main

import (
	"fmt"
	"realworld-go-gin/src/config"
	"realworld-go-gin/src/handlers"

	"github.com/gin-gonic/gin"
)




func main() {
	cfg := config.LoadConfig()

	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/users/login", handlers.Login)
	}

	port := cfg.Server.Port
	fmt.Printf("\n🚀 Сервер запущен ➡️    http://localhost:%s    ⬅️\n", port)
	r.Run(":" + port)
}