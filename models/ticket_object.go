package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketObject struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Subject string `gorm:"not null" json:"subject"`
	Body    string `gorm:"type:jsonb" json:"body"`
	IsHTML  bool   `gorm:"not null;default:false" json:"is_html"`

	IsDraft bool `gorm:"not null;default:false" json:"is_draft"`
	IsSent  bool `gorm:"not null;default:false" json:"is_sent"`

	FromAddress  string         `gorm:"not null" json:"from_address"`
	ToAddresses  pq.StringArray `gorm:"type:text[]" json:"to_addresses"`
	CCAddresses  pq.StringArray `gorm:"type:text[]" json:"cc_addresses"`
	BCCAddresses pq.StringArray `gorm:"type:text[]" json:"bcc_addresses"`

	TicketID string `gorm:"not null" json:"ticket_id"`

	SentAt time.Time `gorm:"not null" json:"sent_at"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Ticket Ticket `gorm:"foreignKey:TicketID" json:"ticket"`
}
