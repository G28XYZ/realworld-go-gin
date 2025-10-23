package database

import (
	"fmt"
	"realworld-go-gin/internal/infrastructure/config_types"
)

func DSN(cfg_db config_types.Database) (string, config_types.Database) {

	const dsnFormat = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

	dsn := fmt.Sprintf(
		dsnFormat,
		cfg_db.Host,
		cfg_db.User,
		cfg_db.Password,
		cfg_db.Name,
		cfg_db.Port,
	)

	return dsn, cfg_db
}
