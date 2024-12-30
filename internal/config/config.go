package config

import (
	"log"
	"os"
	"vrnvgasu/anti-bruteforce/internal/logger"

	"gopkg.in/yaml.v3"
)

var Cfg *Config

type LogLevel string

type Config struct {
	Logger     `yaml:"logger"`
	GRPSServer `yaml:"grpsServer"`
	Redis      `yaml:"redis"`
	PSQL       `json:"psql"`
	Rate       `json:"rate"`
}

func NewConfig(configFile string) *Config {
	c := Config{}

	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Error parsing config file: %s", err.Error())
	}

	return &c
}

type Rate struct {
	Login     int64 `yaml:"login"`
	Password  int64 `yaml:"password"`
	IP        int64 `yaml:"ip"`
	TimeLimit int64 `yaml:"timeLimit"`
}
type PSQL struct {
	DSN       string `yaml:"dsn"`
	Migration string `json:"migration"`
}

type Logger struct {
	Level logger.LogLevel `yaml:"level"`
}

type GRPSServer struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	DB       int    `yaml:"db"`
	TokensDB int    `yaml:"tokenDb"`
}
