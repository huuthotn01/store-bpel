package goods_service

import "time"

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type GetWarehouseResponseData struct {
	GoodsCode   string
	GoodsSize   string
	GoodsColor  string
	WhCode      string
	Quantity    int
	CreatedDate time.Time
	UpdatedDate time.Time
}
