package config

import (
	"fmt"
	"log"
	"os"
	"strings"

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

	Jwt struct {
		Phrase string `yaml:"phrase"`
	} `yaml:"jwt"`
}



func LoadConfig() *Config {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev" // значение по умолчанию
	}
	
	path := fmt.Sprintf("src/config/config.%s.yaml", strings.TrimSpace(env))
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("❌ Ошибка открытия %s: %v", path, err)
	}
	defer file.Close()
	
	decoder := yaml.NewDecoder(file)
	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("❌ Ошибка чтения конфигурации: %v", err)
	}
	
	log.Printf("✅ Конфигурация загружена: %s", path)
	return &cfg
}


func GetConfig() *Config {
	cfg := LoadConfig()
	return cfg
}