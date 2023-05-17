package models

// Wiki represents a wiki entry in the database

type Wikis struct {
	ID          uint   `gorm:"primaryKey"`
	Topic       string `gorm:"not null"`
	Description string `gorm:"default:null"`
	CreatedAt   string `gorm:"column:created_at;not null"`
	UpdatedAt   string `gorm:"column:updated_at;not null"`
}
