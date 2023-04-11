package schema

type UploadImageRequest struct {
	GoodsId    string
	GoodsColor string
	Url        string
	IsDefault  bool
}

type GetGoodsDefaultRequest struct {
	PageNumber int
	PageSize   int
}

type UpdateGoodsRequest struct {
	GoodsSize    string
	GoodsColor   string
	GoodsName    string
	GoodsType    string
	GoodsGender  int
	GoodsAge     string
	Manufacturer string
	UnitPrice    int
	UnitCost     int
	Description  string
}

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
	UnitCost     int
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

type CheckWarehouseRequest struct {
	Elements []*CheckWarehouseRequestElement
}

type CheckWarehouseRequestElement struct {
	GoodsCode  string
	GoodsColor string
	GoodsSize  string
	Quantity   int
}
