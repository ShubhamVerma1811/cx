package model

import (
	"time"
)

type Link struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid"`
	URL       string    `json:"url" gorm:"not null"`
	ShortURL  string    `json:"short_url" gorm:"not null;unique"`
	UserID    uint64    `json:"user_id" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
