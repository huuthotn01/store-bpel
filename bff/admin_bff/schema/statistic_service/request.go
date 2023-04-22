package statistic_service

type GetOverallStatRequest struct {
	Start    string
	End      string
	BranchId []string
}

type FilterGetStatisticRequest struct {
	BranchId []string
	Gender   []int
	Type     []string
	Start    string
	End      string
}

type GetStatOneGoodsRequest struct {
	GoodsId string
	Start   string
	End     string
}

type AddOrderDataRequest struct {
	OrderId         string
	TransactionDate string
	ShopCode        string
	GoodsData       []*AddOrderDataRequest_GoodsData
}

type AddOrderDataRequest_GoodsData struct {
	GoodsId     string
	GoodsSize   string
	GoodsColor  string
	GoodsType   string
	GoodsGender int
	GoodsCost   int
	UnitPrice   int
	Quantity    int
}
