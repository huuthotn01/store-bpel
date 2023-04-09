package statistic_service

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type GetOverallStatisticResponseData struct {
	Revenue int
	Profit  int
	Date    string
}

type GetRevenueResponseData struct {
	Revenue int
	Date    string
}

type GetProfitResponseData struct {
	Profit int
	Date   string
}
