package schema

import (
	"time"
)

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
	Role        int
	IsActivated bool
	CreatedAt   time.Time
}
