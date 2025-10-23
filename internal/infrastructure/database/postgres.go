package database

import (
	"fmt"
	"log"
	"realworld-go-gin/internal/infrastructure/config"
	"realworld-go-gin/internal/infrastructure/config_types"

	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type pgConnectParams struct {
	password string
}

func pg_connect(params pgConnectParams) (*sql.DB, error) {
	cfg := config.GetConfig()

	systemConnStr, _ := DSN(
		config_types.Database{
			Password: params.password,
			Host    : cfg.Database.Host,
			User    : cfg.Database.User,
			Port    : cfg.Database.Port,
			Name    : cfg.Database.Name,
		},
	)

		
	fmt.Println(params.password)
	db, err := sql.Open("postgres", systemConnStr)
	if err != nil {
		// log.Fatalf("Ошибка подключения: %v", err)
		return db, err
	}

	if err := db.Ping(); err != nil {
		// log.Fatalf("❌ Не удалось подключиться к Postgres: %v", err)
		return db, err
	}
	fmt.Println("✅ Подключение к Postgres успешно!")

	return db, nil
}

func Connect() (*gorm.DB, error) {
	cfg := config.GetConfig()

	dsn, db_cfg := DSN(
		config_types.Database{
			Password: cfg.Database.Password,
			Host    : cfg.Database.Host,
			User    : cfg.Database.User,
			Port    : cfg.Database.Port,
			Name    : cfg.Database.Name,
		},
	)

	db, err := pg_connect(pgConnectParams{password: cfg.Database.Password})

	if err != nil {
		db, err = pg_connect(pgConnectParams{password: cfg.Database.AltPass})
		if err != nil {
			log.Fatalf("❌ Не удалось подключиться к Postgres: %v", err)
		} else {
			db_cfg.Password = db_cfg.AltPass
			dsn, db_cfg = DSN(
				config_types.Database{
					Password: cfg.Database.AltPass,
					Host    : cfg.Database.Host,
					User    : cfg.Database.User,
					Port    : cfg.Database.Port,
					Name    : cfg.Database.Name,
				},
			)
		}
	}
	defer db.Close()

	targetDB := db_cfg.Name

	// Проверяем, существует ли нужная БД
	var exists bool
	err = db.QueryRow(`SELECT EXISTS(SELECT datname FROM pg_database WHERE datname = $1)`, targetDB).Scan(&exists)
	if err != nil {
		log.Fatalf("Ошибка проверки БД: %v", err)
	}

	if !exists {
		// Создаём базу
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", targetDB))
		if err != nil {
			log.Fatalf("Ошибка создания БД '%s': %v", targetDB, err)
		}
		fmt.Printf("🎉 База данных '%s' успешно создана!\n", targetDB)
	} else {
		fmt.Printf("ℹ️ База данных '%s' уже существует.\n", targetDB)
	}

	db_gorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к БД: %w", err)
	}

	sqlDB, err := db_gorm.DB()
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении *sql.DB: %w", err)
	}

	// Проверим подключение
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("БД недоступна: %w", err)
	}

	log.Println("✅ Успешное подключение к БД PostgreSQL")
	return db_gorm, nil

}
