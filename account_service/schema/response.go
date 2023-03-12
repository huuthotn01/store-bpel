package schema

import (
	"time"
)

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type SignInResponse struct {
	StatusCode int
	Message    string
	Data       *SignInResponseData
}

type SignInResponseData struct {
	Role int
}

type GetListAccountResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type GetListAccountResponseData struct {
	Username    string
	Role        int
	IsActivated bool
	CreatedAt   time.Time
}
