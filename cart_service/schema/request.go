package schema

type AddCartRequest struct {
	CustomerId string
}

type AddGoodsRequest struct {
	GoodsId    string
	GoodsColor string
	GoodsSize  string
	Quantity   int
}

type DeleteGoodsRequest struct {
	GoodsId    string
	GoodsColor string
	GoodsSize  string
	Quantity   int
}

type UpdateGoodsRequest struct {
	GoodsId    string
	GoodsColor string
	GoodsSize  string
	Quantity   int
}
