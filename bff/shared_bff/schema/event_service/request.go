package event_service

import "encoding/xml"

type GetEventDetailRequest struct {
	XMLName xml.Name `xml:"Body"`
	EventId string
}
