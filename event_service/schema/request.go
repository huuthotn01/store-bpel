package schema

import (
	"time"
)

type AddEventRequest struct {
	Name      string
	Discount  float32
	StartTime time.Time
	EndTime   time.Time
	Image     string
}

type UpdateEventRequest struct {
	Name      string
	Discount  float32
	StartTime time.Time
	EndTime   time.Time
	Image     string
}
