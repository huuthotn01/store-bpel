package event_service

type GetEventDetailRequest struct {
	EventId string
}

type GetEventCurrentRequest struct {
	NextDate string
}

type GetEventByGoodsRequest struct {
	GoodsId string
}
