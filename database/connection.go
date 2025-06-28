package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type DatabaseConnection struct {
	Config DatabaseConfig
	DB     *gorm.DB
}

func NewDatabaseConnectionWithConfig(config DatabaseConfig) *DatabaseConnection {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &DatabaseConnection{Config: config, DB: db}
}

func NewDatabaseConnectionWithEnv() *DatabaseConnection {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	return NewDatabaseConnection(host, port, user, password, dbName, sslMode)
}

func NewDatabaseConnection(host string, port string, user string, password string, dbName string, sslMode string) *DatabaseConnection {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &DatabaseConnection{Config: DatabaseConfig{Host: host, Port: port, User: user, Password: password, DBName: dbName}, DB: db}
}

func (db *DatabaseConnection) AutoMigrate(models ...interface{}) {
	db.DB.AutoMigrate(models...)
}
