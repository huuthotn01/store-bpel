package schema

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetCartResponse struct {
	StatusCode int
	Message    string
	Data       *CartData
}

type CartData struct {
	CartId int
	Goods  []*GoodsData
}

type GoodsData struct {
	GoodsId    string
	GoodsColor string
	GoodsSize  string
	Quantity   int
}
