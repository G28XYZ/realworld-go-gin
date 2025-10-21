package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port  string `yaml:"port"`
		Debug bool   `yaml:"debug"`
	} `yaml:"server"`

	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"database"`
}

func LoadConfig() *Config {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev" // значение по умолчанию
	}

	path := fmt.Sprintf("src/config/config.%s.yaml", env)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("❌ Ошибка открытия %s: %v", path, err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("❌ Ошибка чтения конфигурации: %v", err)
	}

	log.Printf("✅ Конфигурация загружена: %s", path)
	return &cfg
}
