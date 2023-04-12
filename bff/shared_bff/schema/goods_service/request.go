package goods_service

type GetGoodsDefaultRequest struct {
	PageNumber int
	PageSize   int
}

type GetProductsDetailRequest struct {
	GoodsId string
}

type CheckWarehouseRequest struct {
	Elements []*CheckWarehouseRequestElement
}

type CheckWarehouseRequestElement struct {
	GoodsCode  string
	GoodsColor string
	GoodsSize  string
	Quantity   int
}

type CreateGoodsTransactionRequest struct {
	GoodsCode  string
	GoodsColor string
	GoodsSize  string
	Quantity   int
	From       string
	To         string
}
