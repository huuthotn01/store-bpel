package goods_service

import "time"

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
	GoodsId      string
	GoodsName    string
	Classify     []*GetGoodsResponseData_Classify
	GoodsType    string
	GoodsGender  int
	GoodsAge     string
	Manufacturer string
	IsForSale    int
	UnitPrice    int
	UnitCost     int
	Description  string
	Image        []string
}

type GetGoodsResponseData_Classify struct {
	Size  string
	Color string
}

type GetWarehouseResponseData struct {
	GoodsCode   string
	GoodsSize   string
	GoodsColor  string
	WhCode      string
	Quantity    int
	CreatedDate time.Time
	UpdatedDate time.Time
}
