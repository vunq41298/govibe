package model

import "time"

// Province represents a place entity used in business logic or API responses.
// It may contain additional or transformed fields compared to the DB model.
type Province struct {
	ID        int64
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
