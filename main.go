package main

import (
	"fmt"
	"log"
	"realworld-go-gin/internal/application"
	"realworld-go-gin/internal/domain/user"
	"realworld-go-gin/internal/infrastructure/config"
	"realworld-go-gin/internal/infrastructure/database"
	"realworld-go-gin/internal/infrastructure/db"
	"realworld-go-gin/internal/infrastructure/jwt"
	httpHandler "realworld-go-gin/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)


func main() {
	cfg := config.GetConfig()

	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	dbConn, _ := database.Connect()

	// ✅ создаёт таблицу users, если нет
	if err := dbConn.AutoMigrate(&user.User{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}

	repo       := db.NewGormUserRepository(dbConn)
	service    := user.NewService(repo)
	jwtService := jwt.NewService(cfg.Jwt.Phrase)
	usecase    := application.NewUserUseCase(service, jwtService)

	r := gin.Default()
	httpHandler.NewUserHandler(r, usecase)

	fmt.Printf("\nhttp://localhost:%d\n", cfg.Server.Port)
	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}