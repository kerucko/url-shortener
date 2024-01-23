package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server `yaml:"server"`
	Redis  `yaml:"redis"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Redis struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("config path is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var config Config
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatalf("config not read: %v", err)
	}
	return config
}
