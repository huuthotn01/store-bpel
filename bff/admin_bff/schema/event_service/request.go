package event_service

type AddEventRequest struct {
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
	Goods     []string
}

type UpdateEventRequest struct {
	EventId   string
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
	Goods     []string
}

type DeleteEventRequest struct {
	EventId string
}

type UploadImageRequest struct {
	EventId  int
	ImageUrl string
}

type DeleteImageRequest struct {
	EventId int
}
