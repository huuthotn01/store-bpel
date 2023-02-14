package schema

type GetGoodsResponse struct {
	StatusCode int
	Message string
	Result []*GoodsModel
}

type GoodsModel struct {
	GoodsCode string
	GoodsSize string
	GoodsColor string
	GoodsName string
	GoodsType string
	GoodsGender int
	GoodsAge string
	Manufacturer string
	IsForSale bool
	UnitPrice int
	Description string
}

type GetGoodsInWarehouseResponse struct {
	StatusCode int
	Message string
	Result []*GoodsInWarehouseModel
}

type GoodsInWarehouseModel struct {

}

type UpdateResponse struct {
	StatusCode int
	Message string
}
