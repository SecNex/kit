package models

import "time"

type HTTPLog struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	IPAddress string        `gorm:"type:varchar(255)" json:"ip_address"`
	Method    string        `gorm:"type:varchar(255)" json:"method"`
	Path      string        `gorm:"type:varchar(255)" json:"path"`
	Proto     string        `gorm:"type:varchar(255)" json:"proto"`
	Duration  time.Duration `gorm:"type:bigint" json:"duration"`
	UserAgent string        `gorm:"type:varchar(255)" json:"user_agent"`
	Status    int           `gorm:"type:int" json:"status"`
	CreatedAt time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}
