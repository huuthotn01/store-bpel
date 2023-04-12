package schema

import "time"

type GetGoodsDefaultResponse struct {
	StatusCode int
	Message    string
	Data       []*GetGoodsDefaultResponseData
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

type GetDetailProductsResponse struct {
	StatusCode int
	Message    string
	Data       *GetGoodsDefaultResponseData
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

type GetGoodsDetailResponse struct {
	StatusCode int
	Message    string
	Data       *GetGoodsResponseData
}

type GetWarehouseByGoodsResponse struct {
	StatusCode int
	Message    string
	Data       []*GetGoodsInWarehouseResponseData
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
	GoodsCode   string
	GoodsSize   string
	GoodsColor  string
	WhCode      string
	Quantity    int
	CreatedDate time.Time
	UpdatedDate time.Time
}
