package database

import (
	"fmt"
	"log"

	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func Connect() (*gorm.DB, error) {
	dsn, db_cfg := DSN()

	host     := "localhost"
	port     := 5432
	user     := "postgres"
	password := "postgres"
	targetDB := db_cfg.Name

	// 🔹 Подключаемся к системной БД postgres
	systemConnStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password)

	db, err := sql.Open("postgres", systemConnStr)
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Не удалось подключиться к Postgres: %v", err)
	}
	fmt.Println("✅ Подключение к Postgres успешно!")

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
