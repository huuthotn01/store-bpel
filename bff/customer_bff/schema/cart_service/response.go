package cart_service

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type CartData struct {
	CartId int
	Goods  []*GoodsData
}

type GoodsData struct {
	GoodsId      string
	Name         string
	UnitPrice    int
	Price        int
	Images       []string
	ListQuantity []*QuantityData
	Discount     int
	GoodsType    string
	GoodsGender  int
	GoodsAge     string
	Description  string
}

type QuantityData struct {
	GoodsSize  string
	GoodsColor string
	Quantity   int
}
