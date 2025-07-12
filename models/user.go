package models

import (
	"time"

	"github.com/secnex/kit/utils"
	"gorm.io/gorm"
)

type User struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Username string `gorm:"not null;unique" json:"username"`
	Email    string `gorm:"not null;unique" json:"email"`

	FirstName   string `gorm:"not null" json:"first_name"`
	LastName    string `gorm:"not null" json:"last_name"`
	DisplayName string `gorm:"not null" json:"display_name"`

	AvatarURL string `gorm:"not null" json:"avatar_url"`

	Password string `gorm:"not null" json:"password"`

	TenantID string `gorm:"not null" json:"tenant_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Tenant Tenant `gorm:"foreignKey:TenantID" json:"tenant"`

	Teams              []Team              `gorm:"many2many:team_users;" json:"teams"`
	Contacts           []Contact           `gorm:"foreignKey:UserID" json:"contacts"`
	TicketsResolved    []Ticket            `gorm:"foreignKey:ResolvedByUserID" json:"tickets_resolved"`
	RefreshTokens      []RefreshToken      `gorm:"foreignKey:UserID" json:"refresh_tokens"`
	Sessions           []Session           `gorm:"foreignKey:UserID" json:"sessions"`
	AuthorizationCodes []AuthorizationCode `gorm:"foreignKey:UserID" json:"authorization_codes"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := utils.Hash(u.Password, utils.DefaultParams)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return
}
