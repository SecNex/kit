package models

import (
	"time"

	"github.com/secnex/kit/utils"
	"gorm.io/gorm"
)

type Client struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Slug        string `gorm:"not null" json:"slug"`

	ClientSecret string `gorm:"not null" json:"client_secret"`

	ApplicationID string `gorm:"not null" json:"application_id"`

	IsActive bool `gorm:"not null;default:true" json:"is_active"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Application Application `gorm:"foreignKey:ApplicationID" json:"application"`

	AuthorizationCodes []AuthorizationCode `gorm:"foreignKey:ClientID" json:"authorization_codes"`
}

func (c *Client) TableName() string {
	return "clients"
}

func (c *Client) BeforeCreate(tx *gorm.DB) (err error) {
	hashedSecret, err := utils.Hash(c.ClientSecret, utils.DefaultParams)
	if err != nil {
		return err
	}
	c.ClientSecret = hashedSecret
	return
}
