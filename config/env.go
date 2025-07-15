package config

import (
	"os"
	"strconv"

	"github.com/secnex/kit/database"
)

const (
	EnvEnvironment             = "ENV"
	EnvPort                    = "PORT"
	EnvDatabaseHost            = "DATABASE_HOST"
	EnvDatabasePort            = "DATABASE_PORT"
	EnvDatabaseUser            = "DATABASE_USER"
	EnvDatabasePassword        = "DATABASE_PASSWORD"
	EnvDatabaseName            = "DATABASE_NAME"
	EnvDatabaseSSLMode         = "DATABASE_SSL_MODE"
	EnvJWTSecret               = "JWT_SECRET"
	EnvHTTPLog                 = "HTTP_LOG"
	EnvHTTPLogFile             = "HTTP_LOG_FILE"
	EnvDefaultOrganizationName = "DEFAULT_ORGANIZATION_NAME"
	EnvDefaultTenantName       = "DEFAULT_TENANT_NAME"
	EnvDefaultDomainName       = "DEFAULT_DOMAIN_NAME"
)

type DefaultConfig struct {
	OrganizationName string `json:"organization_name"`
	TenantName       string `json:"tenant_name"`
	DomainName       string `json:"domain_name"`
}

type Config struct {
	Environment   string                  `json:"environment"`
	Port          int                     `json:"port"`
	Database      database.DatabaseConfig `json:"database"`
	JWTSecret     string                  `json:"jwt_secret"`
	HTTPLog       bool                    `json:"http_log"`
	HTTPLogFile   string                  `json:"http_log_file"`
	DefaultConfig DefaultConfig           `json:"default_config"`
}

func NewConfig(databaseConfig database.DatabaseConfig, port string, jwtSecret string, httpLog bool, httpLogFile string, defaultConfig DefaultConfig) *Config {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		portInt = 8080
	}

	environment := "development"
	environmentVariable := os.Getenv(EnvEnvironment)
	if environmentVariable != "" {
		environment = environmentVariable
	}

	return &Config{
		Environment:   environment,
		Port:          portInt,
		Database:      databaseConfig,
		JWTSecret:     jwtSecret,
		HTTPLog:       httpLog,
		HTTPLogFile:   httpLogFile,
		DefaultConfig: defaultConfig,
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
		DefaultConfig{
			OrganizationName: os.Getenv(EnvDefaultOrganizationName),
			TenantName:       os.Getenv(EnvDefaultTenantName),
			DomainName:       os.Getenv(EnvDefaultDomainName),
		},
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
