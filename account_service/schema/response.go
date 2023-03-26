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
	Street      string
	Ward        string
	District    string
	Province    string
	Name        string
	IsActivated bool
	CreatedAt   time.Time
}
