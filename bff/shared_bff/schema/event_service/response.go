package event_service

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type GetEventData struct {
	Id        string
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
	Goods     []string
}

type GetEventByGoodsData struct {
	Id        string
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
}
