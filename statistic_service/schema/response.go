package schema

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetOverallStatisticResponse struct {
	StatusCode int
	Message    string
	Data       []*GetOverallStatisticResponseData
}

type GetOverallStatisticResponseData struct {
	Revenue int
	Profit  int
	Date    string
}

type GetRevenueResponse struct {
	StatusCode int
	Message    string
	Data       []*GetRevenueResponseData
}

type GetRevenueOneGoodsResponse struct {
	StatusCode int
	Message    string
	Data       []*GetRevenueResponseData
}

type GetRevenueResponseData struct {
	Revenue int
	Date    string
}

type GetProfitResponse struct {
	StatusCode int
	Message    string
	Data       []*GetProfitResponseData
}

type GetProfitOneGoodsResponse struct {
	StatusCode int
	Message    string
	Data       []*GetProfitResponseData
}

type GetProfitResponseData struct {
	Profit int
	Date   string
}
