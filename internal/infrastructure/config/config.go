package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"realworld-go-gin/internal/infrastructure/config_types"

	"gopkg.in/yaml.v3"
)


func LoadConfig() *config_types.Config {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev" // значение по умолчанию
	}
	
	path := fmt.Sprintf("internal/infrastructure/config/config.%s.yaml", strings.TrimSpace(env))
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("❌ Ошибка открытия %s: %v", path, err)
	}
	defer file.Close()
	
	decoder := yaml.NewDecoder(file)
	var cfg config_types.Config
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("❌ Ошибка чтения конфигурации: %v", err)
	}
	
	log.Printf("✅ Конфигурация загружена: %s", path)
	return &cfg
}


func GetConfig() *config_types.Config {
	cfg := LoadConfig()
	return cfg
}