package goods_service

type AddGoodsRequest struct {
	Element []*AddGoodsRequestData
}

type UpdateGoodsRequest struct {
	Element []*UpdateGoodsRequestData
}

type UpdateGoodsRequestData struct {
	GoodsCode    string
	GoodsSize    string
	GoodsColor   string
	GoodsName    string
	GoodsType    string
	GoodsGender  int
	GoodsAge     string
	Manufacturer string
	IsForSale    bool
	UnitPrice    int
	UnitCost     int
	Description  string
}

type AddGoodsRequestData struct {
	GoodsSize    string
	GoodsColor   string
	GoodsName    string
	GoodsType    string
	GoodsGender  int
	GoodsAge     string
	Manufacturer string
	IsForSale    bool
	UnitPrice    int
	Description  string
}

type CreateGoodsTransactionRequest struct {
	GoodsCode  string
	GoodsColor string
	GoodsSize  string
	Quantity   int
	From       string
	To         string
}

type GetWarehouseByGoodsRequest struct {
	GoodsCode string
}
