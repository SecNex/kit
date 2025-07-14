package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkQueuesEntryExternalType string

type WorkQueuesEntryStatus string

const (
	WorkQueuesEntryExternalTypeRabbitMQ WorkQueuesEntryExternalType = "rabbitmq"

	WorkQueuesEntryStatusNew        WorkQueuesEntryStatus = "new"
	WorkQueuesEntryStatusPending    WorkQueuesEntryStatus = "pending"
	WorkQueuesEntryStatusProcessing WorkQueuesEntryStatus = "processing"
	WorkQueuesEntryStatusCompleted  WorkQueuesEntryStatus = "completed"
	WorkQueuesEntryStatusFailed     WorkQueuesEntryStatus = "failed"
)

type WorkQueuesEntry struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	WorkQueueID string `gorm:"not null" json:"work_queue_id"`

	ExternalID   string                      `gorm:"not null" json:"external_id"`
	ExternalType WorkQueuesEntryExternalType `gorm:"not null" json:"external_type"`

	Status WorkQueuesEntryStatus `gorm:"not null" json:"status"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	WorkQueue WorkQueue `gorm:"foreignKey:WorkQueueID" json:"work_queue"`
}
