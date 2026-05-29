package models

import "time"

type OAuthState struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time
}