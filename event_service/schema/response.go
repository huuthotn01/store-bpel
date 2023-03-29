package schema

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetEventResponse struct {
	StatusCode int
	Message    string
	Data       []*GetEventData
}

type GetEventDetailResponse struct {
	StatusCode int
	Message    string
	Data       *GetEventData
}

type GetEventByGoodsResponse struct {
	StatusCode int
	Message    string
	Data       []*GetEventByGoodsData
}

type GetEventData struct {
	Id        int
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
	Goods     []string
}

type GetEventByGoodsData struct {
	Id        int
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
}
