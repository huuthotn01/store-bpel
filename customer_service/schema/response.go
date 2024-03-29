package schema

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetCustomerInfoResponse struct {
	StatusCode int
	Message    string
	Data       *GetCustomerInfoResponseData
}

type GetCustomerInfoResponseData struct {
	Username string
	Email    string
	Name     string
	Phone    string
	Gender   string
	Age      int32
	Street   string
	Ward     string
	District string
	Province string
}
