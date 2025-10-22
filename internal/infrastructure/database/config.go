package database

import (
	"fmt"
	"realworld-go-gin/internal/infrastructure/config"
	"realworld-go-gin/internal/infrastructure/config_types"
)

func DSN() (string, config_types.Database) {
	cfg := config.GetConfig()

	const dsnFormat = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

	dsn := fmt.Sprintf(
		dsnFormat,
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)

	return dsn, cfg.Database
}
