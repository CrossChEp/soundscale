package config

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Service struct {
	Port    int    `yaml:"port"`
	Address string `yaml:"address"`
}


type Config struct {
	Port     int                `yaml:"port"`
	Address  string             `yaml:"address"`
	Services map[string]Service `yaml:"services"`
}

var (
	CONFIG          Config        = Config{}
	EXPIRATION_TIME time.Duration = time.Minute * 30
	Logger                        = log.New()
)
