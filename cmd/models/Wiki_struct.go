package models

import "time"

// Wiki represents a wiki entry in the database

type Wikis struct {
	ID          uint      `gorm:"primaryKey"`
	Topic       string    `gorm:"not null"`
	Description string    `gorm:"default:null"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null"`
}
