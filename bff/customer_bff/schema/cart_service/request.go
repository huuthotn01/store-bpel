package cart_service

type GetCartRequest struct {
	UserId string
}

type AddGoodsRequest struct {
	CartId string
	Goods  []*GoodsRequestData
}

type UpdateGoodsRequest struct {
	CartId string
	Goods  []*GoodsRequestData
}

type DeleteGoodsRequest struct {
	CartId string
	Goods  []*DeleteGoodsData
}

type DeleteAllGoodsRequest struct {
	CartId string
}

type DeleteGoodsData struct {
	GoodsId    string
	GoodsColor string
	GoodsSize  string
}

type GoodsRequestData struct {
	GoodsId    string
	GoodsColor string
	GoodsSize  string
	Quantity   int
}
