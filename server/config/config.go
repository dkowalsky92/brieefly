package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"
)

type contextKey string

var (
	configKey contextKey
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

	configProductionFile  string = "config-prod.json"
	configDevelopmentFile string = "config-dev.json"
	configLocalFile       string = "config-local.json"
	// ConfigFilePath - path to config files
	ConfigFilePath string = "./secrets/"
)

// Config - stores all necessary information regarding server & database setup
type Config struct {
	Environment Environment          `json:"-"`
	Database    *DatabaseParams      `json:"database"`
	Server      *ServerParams        `json:"server"`
	Auth        *AuthorizationParams `json:"auth"`
}

// DatabaseParams - database info
type DatabaseParams struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
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

// AuthorizationParams - jwt authorization info
type AuthorizationParams struct {
	Public  string `json:"public"`
	Private string `json:"private"`
}

// NewConfig - creates a new Config object with specified environment
func NewConfig(environment Environment) (*Config, *err.Error) {
	var file *os.File
	var cnfErr error

	switch environment {
	case Production:
		file, cnfErr = os.Open(fmt.Sprintf("%s%s", ConfigFilePath, configProductionFile))
	case Development:
		file, cnfErr = os.Open(fmt.Sprintf("%s%s", ConfigFilePath, configDevelopmentFile))
	case Local:
		file, cnfErr = os.Open(fmt.Sprintf("%s%s", ConfigFilePath, configLocalFile))
	}

	if cnfErr != nil {
		log.Error(cnfErr)
		return nil, err.New(cnfErr, err.ErrConfigMalformed, nil)
	}

	var config Config
	decoder := json.NewDecoder(file)
	cnfErr = decoder.Decode(&config)
	if cnfErr != nil {
		log.Error(cnfErr)
		return nil, err.New(cnfErr, err.ErrConfigMalformed, nil)
	}

	config.Environment = environment

	return &config, nil
}

// TLSCert - get public TLS key
func (c *Config) TLSCert() string {
	return fmt.Sprintf("%s%s", ConfigFilePath, c.Server.Certificate)
}

// TLSKey - get private TLS key
func (c *Config) TLSKey() string {
	return fmt.Sprintf("%s%s", ConfigFilePath, c.Server.Key)
}

// IntoContext - inserts the associated config into given context
func IntoContext(ctx context.Context, config *Config) context.Context {
	return context.WithValue(ctx, configKey, config)
}

// FromContext - returns the associated config from given context
func FromContext(ctx context.Context) *Config {
	return ctx.Value(configKey).(*Config)
}

// MyPath - generates a path with parameters from given config
func MyPath(config *Config) string {
	address := fmt.Sprintf("%s:%s", config.Server.IP, config.Server.Port)
	return address
}
