package goods_service

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
	GoodsCode  string
	GoodsSize  string
	GoodsColor string
	WhCode     string
	Quantity   int
}
