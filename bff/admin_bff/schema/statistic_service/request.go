package statistic_service

type GetOverallStatRequest struct {
	Start string
	End   string
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
