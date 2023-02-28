package schema

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type UpdateResponse struct {
	StatusCode int
	Message    string
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
	Description  string
}

type GetGoodsInWarehouseResponseData struct {
}
