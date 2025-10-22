package config_types

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type Server struct {
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

type Jwt struct {
	Phrase string `yaml:"phrase"`
}

type Config struct {
	Server Server `yaml:"server"`

	Database Database `yaml:"database"`

	Jwt Jwt `yaml:"jwt"`
}