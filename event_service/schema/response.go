package schema

import "time"

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetEventResponse struct {
	StatusCode int
	Message    string
	Data       []*GetEventData
}

type GetEventData struct {
	Id        string
	Name      string
	Discount  float32
	StartTime time.Time
	EndTime   time.Time
	Image     string
	Goods     []*string
}
