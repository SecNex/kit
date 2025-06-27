package models

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`

	IsActive  bool `gorm:"not null;default:true" json:"is_active"`
	IsDefault bool `gorm:"not null;default:false" json:"is_default"`

	TenantID string `gorm:"not null" json:"tenant_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Tenant Tenant `gorm:"foreignKey:TenantID" json:"tenant"`

	Users []User `gorm:"many2many:team_users;" json:"users"`
}

func (t *Team) TableName() string {
	return "teams"
}
