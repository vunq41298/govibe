package model

import "time"

// Image represents a place entity used in business logic or API responses.
// It may contain additional or transformed fields compared to the DB model.
type Image struct {
	ID            int64
	PlaceID       int64
	URL           string
	Type          string
	Caption       string
	OrderNum      int
	IsHeaderImage bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
