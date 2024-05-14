package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"haha/internal/logger"
)

var CONFIG Config

type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		URL string `yaml:"url"`
	} `yaml:"server"`
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     int64  `yaml:"port"`
		Path     string `yaml:"path"`
	} `yaml:"db"`
}

func InitConfig(logg logger.Logger) (Config, error) {
	if err := cleanenv.ReadConfig("./config/local.yml", &CONFIG); err != nil {
		logg.Fatal(err)
		return Config{}, err
	}

	logg.Info("init config")

	return CONFIG, nil
}

func GetConf() Config {
	return CONFIG
}
