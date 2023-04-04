package goods_service

type AddGoodsRequest struct {
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
