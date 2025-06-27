package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Application struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Slug        string `gorm:"not null" json:"slug"`

	TenantID string `gorm:"not null" json:"tenant_id"`

	IsActive     bool           `gorm:"not null;default:true" json:"is_active"`
	RedirectURIs pq.StringArray `gorm:"type:text[]" json:"redirect_uris"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Tenant Tenant `gorm:"foreignKey:TenantID" json:"tenant"`

	Client []Client `gorm:"foreignKey:ApplicationID" json:"clients"`
}

func (a *Application) TableName() string {
	return "applications"
}
