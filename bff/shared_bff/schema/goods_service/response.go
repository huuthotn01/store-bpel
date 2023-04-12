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

type GetGoodsDefaultResponseData struct {
	GoodsId      string
	Name         string
	UnitPrice    int
	Price        int
	Images       []string
	ListQuantity []*GetGoodsDefault_QuantityList
	Discount     int
	GoodsType    string
	GoodsGender  int
	GoodsAge     string
	Description  string
}

type GetGoodsDefault_QuantityList struct {
	GoodsSize  string
	GoodsColor string
	Quantity   int
}

type CheckWarehouseResponseData struct {
	NeedTransfer     bool
	WarehouseActions []*WarehouseActions
}

type WarehouseActions struct {
	GoodsCode  string
	GoodsColor string
	GoodsSize  string
	Quantity   int
	From       string
	To         string
}
