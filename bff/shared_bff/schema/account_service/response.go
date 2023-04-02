package account_service

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type SignInResponseData struct {
	UserId string
	Role   int
	Token  string
}
