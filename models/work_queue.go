package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkQueue struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`

	WorkerID string `gorm:"not null" json:"worker_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Worker Worker `gorm:"foreignKey:WorkerID" json:"worker"`

	Entries []WorkQueuesEntry `gorm:"foreignKey:WorkQueueID" json:"entries"`
}

func (w *WorkQueue) TableName() string {
	return "work_queues"
}
