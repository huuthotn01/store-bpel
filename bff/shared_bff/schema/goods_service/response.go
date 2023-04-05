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

type GetGoodsResponseData struct {
	GoodsCode    string
	GoodsSize    string
	GoodsColor   string
	GoodsName    string
	GoodsType    string
	GoodsGender  int
	GoodsAge     string
	Manufacturer string
	IsForSale    int
	UnitPrice    int
	UnitCost     int
	Description  string
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