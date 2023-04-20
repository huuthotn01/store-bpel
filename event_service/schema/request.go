package schema

type AddEventRequest struct {
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
	Goods     []string
}

type UpdateEventRequest struct {
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
	Goods     []string
}

type UploadImageRequest struct {
	EventId  string
	ImageUrl string
}
