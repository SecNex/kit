package models

import (
	"time"

	"gorm.io/gorm"
)

type Queue struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Slug        string `gorm:"not null" json:"slug"`

	ProjectID string `gorm:"not null" json:"project_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Project Project `gorm:"foreignKey:ProjectID" json:"project"`
}

func (q *Queue) TableName() string {
	return "queues"
}
