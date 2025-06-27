package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	UserID string `gorm:"not null" json:"user_id"`

	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}

func (s *Session) TableName() string {
	return "sessions"
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	s.ExpiresAt = time.Now().Add(time.Minute * 60)
	return
}
