package models

import "time"

type HTTPLog struct {
	ID        uint          `gorm:"primaryKey"`
	IPAddress string        `gorm:"type:varchar(255)"`
	Method    string        `gorm:"type:varchar(255)"`
	Path      string        `gorm:"type:varchar(255)"`
	Proto     string        `gorm:"type:varchar(255)"`
	Duration  time.Duration `gorm:"type:bigint"`
	UserAgent string        `gorm:"type:varchar(255)"`
	Status    int           `gorm:"type:int"`
	CreatedAt time.Time     `gorm:"autoCreateTime"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime"`
}
