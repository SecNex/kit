package config

import (
	"os"
	"strconv"

	"github.com/secnex/kit/database"
)

const (
	EnvEnvironment      = "ENV"
	EnvPort             = "PORT"
	EnvDatabaseHost     = "DATABASE_HOST"
	EnvDatabasePort     = "DATABASE_PORT"
	EnvDatabaseUser     = "DATABASE_USER"
	EnvDatabasePassword = "DATABASE_PASSWORD"
	EnvDatabaseName     = "DATABASE_NAME"
	EnvDatabaseSSLMode  = "DATABASE_SSL_MODE"
	EnvJWTSecret        = "JWT_SECRET"
	EnvHTTPLog          = "HTTP_LOG"
	EnvHTTPLogFile      = "HTTP_LOG_FILE"
)

type Config struct {
	Environment string
	Port        int
	Database    database.DatabaseConfig
	JWTSecret   string
	HTTPLog     bool
	HTTPLogFile string
}

func NewConfig(databaseConfig database.DatabaseConfig, port string, jwtSecret string, httpLog bool, httpLogFile string) *Config {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		portInt = 8080
	}

	return &Config{
		Environment: "development",
		Port:        portInt,
		Database:    databaseConfig,
		JWTSecret:   jwtSecret,
		HTTPLog:     httpLog,
		HTTPLogFile: httpLogFile,
	}
}

func NewConfigEnv() *Config {
	return NewConfig(
		database.DatabaseConfig{
			Host:     os.Getenv(EnvDatabaseHost),
			Port:     os.Getenv(EnvDatabasePort),
			User:     os.Getenv(EnvDatabaseUser),
			Password: os.Getenv(EnvDatabasePassword),
			DBName:   os.Getenv(EnvDatabaseName),
			SSLMode:  os.Getenv(EnvDatabaseSSLMode),
		},
		os.Getenv(EnvPort),
		os.Getenv(EnvJWTSecret),
		os.Getenv(EnvHTTPLog) == "true",
		os.Getenv(EnvHTTPLogFile),
	)
}

func (c *Config) GetDatabaseConfig() database.DatabaseConfig {
	return c.Database
}

func (c *Config) GetJWTSecret() string {
	return c.JWTSecret
}

func (c *Config) GetHTTPLog() bool {
	return c.HTTPLog
}

func (c *Config) GetHTTPLogFile() string {
	return c.HTTPLogFile
}
