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
	UserId string
	Role   int
	Token  string
}

type GetListAccountResponse struct {
	StatusCode int
	Message    string
	Data       []*GetListAccountResponseData
}

type GetListAccountResponseData struct {
	Username    string
	Id          string
	Role        int
	PhoneNumber string
	Email       string
	Name        string
	IsActivated bool
	CreatedAt   time.Time
}
