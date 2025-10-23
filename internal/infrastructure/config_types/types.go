package config_types

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	AltPass  string `yaml:"alt_pass"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Server struct {
	Port  int  `yaml:"port"`
	Debug bool `yaml:"debug"`
}

type Jwt struct {
	Phrase string `yaml:"phrase"`
}

type Config struct {
	Server Server `yaml:"server"`

	Database Database `yaml:"database"`

	Jwt Jwt `yaml:"jwt"`
}