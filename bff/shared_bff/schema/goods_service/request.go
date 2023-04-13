package goods_service

type SearchGoodsRequest struct {
	Query    string
	PageSize int
	Category int // 0: nothing, 1: best-selling goods, 2: newly added goods
}

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
