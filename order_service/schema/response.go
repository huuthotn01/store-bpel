package schema

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetOnlineOrdersStatusResponse struct {
	StatusCode int
	Message    string
	Data       []*GetOnlineOrdersStatusResponseData
}

type GetOnlineOrdersStatusResponseData struct {
	OrderId   int
	State     string
	StateTime string
}
