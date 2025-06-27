package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null;unique" json:"email"`

	Phone   string `gorm:"default:null" json:"phone"`
	Address string `gorm:"default:null" json:"address"`
	City    string `gorm:"default:null" json:"city"`
	State   string `gorm:"default:null" json:"state"`
	Zip     string `gorm:"default:null" json:"zip"`
	Country string `gorm:"default:null" json:"country"`

	UserID string `gorm:"default:null" json:"user_id"`

	TenantID string `gorm:"not null" json:"tenant_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Tenant Tenant `gorm:"foreignKey:TenantID" json:"tenant"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`

	Tickets []Ticket `gorm:"foreignKey:CreatedByContactID" json:"tickets"`
}

func (c *Contact) TableName() string {
	return "contacts"
}
