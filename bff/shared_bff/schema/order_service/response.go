package order_service

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type GetShipFeeResponseData struct {
	ShipFee      int
	ExpectedDate string
}
