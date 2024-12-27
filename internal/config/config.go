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
	Logger     LoggerConf     `yaml:"logger"`
	GRPSServer GRPSServerConf `yaml:"grpsServer"`
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

type LoggerConf struct {
	Level logger.LogLevel `yaml:"level"`
}

type GRPSServerConf struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}
