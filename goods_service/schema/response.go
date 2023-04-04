package schema

type GetDetailGoodsResponse struct {
	StatusCode int
	Message    string
	Data       *GetGoodsResponseData
}

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetGoodsResponse struct {
	StatusCode int
	Message    string
	Data       []*GetGoodsResponseData
}

type GetWarehouseByGoodsResponse struct {
	StatusCode int
	Message    string
	Data       []*GetGoodsInWarehouseResponseData
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

type CheckWarehouseResponse struct {
	StatusCode int
	Message    string
	Data       *CheckWarehouseResponseData
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

type GetGoodsInWarehouseResponseData struct {
	GoodsCode  string
	GoodsSize  string
	GoodsColor string
	WhCode     string
	Quantity   int
}
