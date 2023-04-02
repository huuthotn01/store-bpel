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

type GetListAccountResponseData struct {
	Username    string
	Id          string
	Role        int
	PhoneNumber string
	StartDate   time.Time
	BirthDate   string
	Street      string
	Ward        string
	District    string
	Province    string
	Name        string
	IsActivated bool
	CreatedAt   time.Time
}
