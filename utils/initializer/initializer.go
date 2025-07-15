package initializer

import (
	"fmt"

	"github.com/secnex/kit/config"
	"github.com/secnex/kit/database"
	"github.com/secnex/kit/logging"
	"github.com/secnex/kit/models"
)

const (
	DEFAULT_ID     = "00000000-0000-0000-0000-000000000000"
	ADMIN_USERNAME = "admin"
	ADMIN_EMAIL    = "admin@example.com"
)

type Initializer struct {
	Database *database.DatabaseConnection
	Config   *config.Config
}

func NewInitializer(database *database.DatabaseConnection, config *config.Config) *Initializer {
	return &Initializer{
		Database: database,
		Config:   config,
	}
}

func (i *Initializer) Initialize() {
	fmt.Printf("🚀 Initializing...\n")

	// Check if database is connected
	i.Database.TestConnection()

	i.DefaultOrganization()

	fmt.Printf("✅ Initialized!\n\n")
}

func (i *Initializer) DefaultOrganization() {
	organization := models.Organization{
		ID:   DEFAULT_ID,
		Name: i.Config.DefaultConfig.OrganizationName,
	}

	// Check if organization already exists
	var existingOrganization models.Organization
	err := i.Database.DB.Where("id = ? or name = ?", DEFAULT_ID, i.Config.DefaultConfig.OrganizationName).First(&existingOrganization).Error
	// If error, create default organization
	if err != nil {
		logging.ErrorWithErr("Failed to check if default organization exists", err)
		logging.Info("Creating default organization")

		var createdOrganization models.Organization
		err = i.Database.DB.Create(&organization).Scan(&createdOrganization).Error
		if err != nil {
			logging.ErrorWithErr("Failed to create default organization", err)
			return
		}

		existingOrganization = createdOrganization
	}

	logging.InfoWithFields("Default organization", map[string]interface{}{
		"organization": existingOrganization,
	})
}

func (i *Initializer) DefaultTenant() {
	tenant := models.Tenant{
		ID:   DEFAULT_ID,
		Name: i.Config.DefaultConfig.TenantName,
	}

	// Check if tenant already exists
	var existingTenant models.Tenant
	err := i.Database.DB.Where("id = ? or name = ?", DEFAULT_ID, i.Config.DefaultConfig.TenantName).First(&existingTenant).Error
	// If error, create default tenant
	if err != nil {
		logging.ErrorWithErr("Failed to check if default tenant exists", err)
		logging.Info("Creating default tenant")

		err = i.Database.DB.Create(&tenant).Error
		if err != nil {
			logging.ErrorWithErr("Failed to create default tenant", err)
			return
		}
	}

	logging.InfoWithFields("Default tenant created", map[string]interface{}{
		"tenant": tenant,
	})
}
