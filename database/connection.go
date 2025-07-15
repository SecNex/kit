package database

import (
	"fmt"
	"os"

	"github.com/secnex/kit/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	return NewDatabaseConnection(config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
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
	fmt.Println("🚀 Connecting to database...")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Printf("🚨 Failed to connect to database...\n\n")
		panic(err)
	}

	fmt.Printf("✅ Connected to database!\n\n")

	return &DatabaseConnection{Config: DatabaseConfig{Host: host, Port: port, User: user, Password: password, DBName: dbName}, DB: db}
}

func (db *DatabaseConnection) AutoMigrateAll() {
	db.AutoMigrate(
		// Base models without dependencies
		&models.HTTPLog{},
		&models.Organization{},
		&models.Domain{},

		// Tenant depends on Organization and Domain
		&models.Tenant{},

		// Models depending on Tenant
		&models.User{},        // depends on Tenant
		&models.Team{},        // depends on Tenant
		&models.Application{}, // depends on Tenant
		&models.Project{},     // depends on Tenant

		// Models depending on User
		&models.Contact{},      // depends on User
		&models.Session{},      // depends on User
		&models.RefreshToken{}, // depends on User

		// Models depending on Application
		&models.Client{}, // depends on Application

		// Models depending on Project
		&models.Queue{}, // depends on Project

		// Models depending on Worker
		&models.Worker{},          // depends on Project
		&models.WorkQueue{},       // depends on Worker
		&models.WorkQueuesEntry{}, // depends on WorkQueue

		// Models depending on multiple entities
		&models.AuthorizationCode{}, // depends on Client and User
		&models.Ticket{},            // depends on Queue, Contact, User, Tenant
		&models.TicketObject{},      // depends on Ticket, Tenant
	)
}

func (db *DatabaseConnection) AutoMigrate(models ...interface{}) {
	fmt.Println("🚀 Migrating models...")
	db.DB.AutoMigrate(models...)
	fmt.Println("✅ Models migrated!")
}

func (db *DatabaseConnection) TestConnection() error {
	fmt.Printf("\n🔄 Testing database connection...\n")
	if err := db.DB.Exec("SELECT 1").Error; err != nil {
		fmt.Printf("🚨 Failed to test database connection!\n\n")
		return err
	}
	fmt.Printf("✅ Database connection tested!\n\n")
	return nil
}
