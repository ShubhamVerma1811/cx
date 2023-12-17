package model

import "time"

type Stats struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	LinkID    uint64    `json:"link_id" gorm:"not null"`
	IP        string    `json:"ip" gorm:"not null"`
	Referer   string    `json:"referer" gorm:"not null"`
	UserAgent string    `json:"user_agent" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Clicks    int       `json:"clicks" gorm:"not null;default:0"`
}
