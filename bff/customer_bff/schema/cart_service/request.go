package cart_service

type GetCartRequest struct {
	UserId string
}

type AddGoodsRequest struct {
	CartId string
	Goods  []AddGoodsRequestData
}

type DeleteGoodsRequest struct {
	GoodsId    string
	GoodsColor string
	GoodsSize  string
}

type UpdateGoodsRequest struct {
	GoodsId    string
	GoodsColor string
	GoodsSize  string
	Quantity   int
}

type AddGoodsRequestData struct {
	GoodsId    string
	GoodsColor string
	GoodsSize  string
	Quantity   int
}
