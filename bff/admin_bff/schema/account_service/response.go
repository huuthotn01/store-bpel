package account_service

import "time"

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
	Role int
}

type GetListAccountResponseData struct {
	Username    string
	Role        int
	IsActivated bool
	CreatedAt   time.Time
}
