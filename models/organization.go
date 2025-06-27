package models

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name string `gorm:"not null;unique" json:"name"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Tenants []Tenant `gorm:"foreignKey:OrganizationID" json:"tenants"`
	Domains []Domain `gorm:"foreignKey:OrganizationID" json:"domains"`
}

func (o *Organization) TableName() string {
	return "organizations"
}
