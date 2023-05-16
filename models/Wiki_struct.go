package models

import "time"

// Wiki represents a wiki entry in the database

type Wiki struct {
	ID          uint      `gorm:"primaryKey"`
	Topic       string    `gorm:"not null"`
	Description string    `gorm:"default:null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
