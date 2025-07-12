package models

import (
	"time"

	"github.com/secnex/kit/utils"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	UserID string `gorm:"not null" json:"user_id"`

	Token string `gorm:"not null" json:"token"`

	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}

func (r *RefreshToken) TableName() string {
	return "refresh_tokens"
}

func (r *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	r.ExpiresAt = time.Now().Add(time.Hour * 24)

	hashedToken, err := utils.Hash(r.Token, utils.DefaultParams)
	if err != nil {
		return err
	}
	r.Token = hashedToken

	return
}
