package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dkowalsky/brieefly/log"
)

// Environment - defines environment
type Environment uint8

const (
	// Production - production environment constant
	Production Environment = 0
	// Development - development environment constant
	Development Environment = 1
	// Local - local environment constant
	Local Environment = 2

	configProductionPath  string = "secret/config-prod.json"
	configDevelopmentPath string = "secret/config-dev.json"
	configLocalPath       string = "secret/config-local.json"
)

// Config - stores all necessary information regarding server & database setup
type Config struct {
	Database *DatabaseParams `json:"database"`
	Server   *ServerParams   `json:"server"`
}

// DatabaseParams - database info
type DatabaseParams struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// ServerParams - server info
type ServerParams struct {
	Certificate  string `json:"certificate"`
	Key          string `json:"key"`
	Protocol     string `json:"protocol"`
	FileProtocol string `json:"file_protocol"`
	IP           string `json:"ip"`
	Port         string `json:"port"`
}

// NewConfig - creates a new Config object with specified environment
func NewConfig(environment Environment) (*Config, error) {
	var file *os.File
	var err error

	switch environment {
	case Production:
		file, err = os.Open(configProductionPath)
	case Development:
		file, err = os.Open(configDevelopmentPath)
	case Local:
		file, err = os.Open(configLocalPath)
	}

	if err != nil {
		log.Error(err)
		return nil, err
	}

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info(config)
	return &config, nil
}

// MyPath - generates a path with parameters from given config
func MyPath(config *Config) string {
	address := fmt.Sprintf("%s:%s", config.Server.IP, config.Server.Port)
	return address
}
