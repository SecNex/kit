package models

import (
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	ID          string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`

	IsActive  bool `gorm:"not null;default:true" json:"is_active"`
	IsDefault bool `gorm:"not null;default:false" json:"is_default"`

	OrganizationID string `gorm:"not null" json:"organization_id"`
	DomainID       string `gorm:"not null" json:"domain_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Organization Organization `gorm:"foreignKey:OrganizationID" json:"organization"`
	Domain       Domain       `gorm:"foreignKey:DomainID" json:"domain"`

	Users    []User        `gorm:"foreignKey:TenantID" json:"users"`
	Teams    []Team        `gorm:"foreignKey:TenantID" json:"teams"`
	Apps     []Application `gorm:"foreignKey:TenantID" json:"apps"`
	Projects []Project     `gorm:"foreignKey:TenantID" json:"projects"`
}

func (t *Tenant) TableName() string {
	return "tenants"
}
