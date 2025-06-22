package model

import (
	"time"
)

// Place represents a place entity used in business logic or API responses.
// It may contain additional or transformed fields compared to the DB model.
type Place struct {
	ID          int64
	Name        string
	Description string
	ProvinceID  int64
	Slug        string
	AverageVote float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
