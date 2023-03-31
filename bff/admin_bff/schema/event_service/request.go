package event_service

import "encoding/xml"

type AddEventRequest struct {
	XMLName   xml.Name `xml:"Body"`
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
	Goods     []string
}

type UpdateEventRequest struct {
	XMLName   xml.Name `xml:"Body"`
	EventId   string
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
	Goods     []string
}

type DeleteEventRequest struct {
	XMLName xml.Name `xml:"Body"`
	EventId string
}
