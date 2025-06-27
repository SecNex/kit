package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	TicketNumber string `gorm:"not null" json:"ticket_number"`

	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`

	QueueID string `gorm:"not null" json:"queue_id"`

	CreatedByContactID string `gorm:"not null" json:"created_by_contact_id"`
	ResolvedByUserID   string `gorm:"default:null" json:"resolved_by_user_id"`

	ResolvedAt time.Time `gorm:"default:null" json:"resolved_at"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Queue      Queue   `gorm:"foreignKey:QueueID" json:"queue"`
	CreatedBy  Contact `gorm:"foreignKey:CreatedByContactID" json:"created_by"`
	ResolvedBy User    `gorm:"foreignKey:ResolvedByUserID" json:"resolved_by"`

	TicketObjects []TicketObject `gorm:"foreignKey:TicketID" json:"ticket_objects"`
}

func (t *Ticket) TableName() string {
	return "tickets"
}
