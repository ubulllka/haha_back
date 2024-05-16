package config

import (
	"haha/internal/logger"
	"os"
)

var CONFIG *Config

//type Config struct {
//	Env    string `yaml:"env"`
//	Server struct {
//		URL string `yaml:"url"`
//	} `yaml:"server"`
//	DB struct {
//		User     string `yaml:"user"`
//		Password string `yaml:"password"`
//		Name     string `yaml:"name"`
//		Host     string `yaml:"host"`
//		Port     int64  `yaml:"port"`
//		Path     string `yaml:"path"`
//	} `yaml:"db"`
//}

type DB struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Host     string `env:"HOST"`
	Port     string `env:"POSR"`
}

type Server struct {
	URL string `env:"SERVER_URL"`
}

type Client struct {
	URL string `env:"CLIENT_URL"`
}

type Config struct {
	Env    string `env:"ENV"`
	DB     DB
	Server Server
	Client Client
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func InitConfig(logg logger.Logger) *Config {
	CONFIG = &Config{
		Env: getEnv("ENV", ""),
		DB: DB{
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", ""),
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
		},
		Server: Server{
			getEnv("SERVER_URL", ""),
		},
		Client: Client{
			URL: getEnv("CLIENT_URL", ""),
		},
	}
	logg.Info("init config")
	return CONFIG
}

func GetConf() *Config {
	return CONFIG
}
