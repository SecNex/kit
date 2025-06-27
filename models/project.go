package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Slug        string `gorm:"not null" json:"slug"`

	TenantID string `gorm:"not null" json:"tenant_id"`

	IsActive bool `gorm:"not null;default:true" json:"is_active"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Tenant Tenant `gorm:"foreignKey:TenantID" json:"tenant"`

	Queues []Queue `gorm:"foreignKey:ProjectID" json:"queues"`
}

func (p *Project) TableName() string {
	return "projects"
}
