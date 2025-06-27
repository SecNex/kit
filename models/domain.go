package models

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name string `gorm:"not null;unique" json:"name"`

	IsActive   bool `gorm:"not null;default:true" json:"is_active"`
	IsVerified bool `gorm:"not null;default:false" json:"is_verified"`
	IsDefault  bool `gorm:"not null;default:false" json:"is_default"`

	OrganizationID string `gorm:"not null" json:"organization_id"`

	VerifiedAt time.Time `gorm:"not null;default:null" json:"verified_at"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Organization Organization `gorm:"foreignKey:OrganizationID" json:"organization"`
}

func (d *Domain) TableName() string {
	return "domains"
}
