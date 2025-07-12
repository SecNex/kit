package models

import (
	"time"

	"github.com/secnex/kit/utils"
	"gorm.io/gorm"
)

type AuthorizationCode struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	ClientID string `gorm:"not null" json:"client_id"`
	UserID   string `gorm:"not null" json:"user_id"`

	Code        string `gorm:"not null" json:"code"`
	RedirectURI string `gorm:"not null" json:"redirect_uri"`

	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Client Client `gorm:"foreignKey:ClientID" json:"client"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}

func (a *AuthorizationCode) TableName() string {
	return "authorization_codes"
}

func (a *AuthorizationCode) BeforeCreate(tx *gorm.DB) (err error) {
	a.ExpiresAt = time.Now().Add(time.Minute * 5)

	hashedCode, err := utils.Hash(a.Code, utils.DefaultParams)
	if err != nil {
		return err
	}
	a.Code = hashedCode

	return
}
