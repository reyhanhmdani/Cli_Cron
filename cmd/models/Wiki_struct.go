package models

import "time"

type Wikis struct {
	ID          int
	Topic       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
