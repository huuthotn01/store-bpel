package staff_service

import (
	"database/sql"
	"time"
)

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type GetStaffResponseData struct {
	StaffId     string
	StaffName   string
	Street      string
	Ward        string
	District    string
	Province    string
	CitizenId   string
	Role        string
	BranchId    string
	Hometown    string
	Salary      int
	Birthdate   string
	StartDate   time.Time
	Gender      string
	PhoneNumber string
	Status      string
	Email       string
}

type GetStaffAttendanceResponseData struct {
	AttendanceDate string
	CheckinTime    time.Time
	CheckoutTime   sql.NullTime
}

type GetRequestResponseData struct {
	Id          string
	RequestDate time.Time
	RequestType string // ADD or DELETE
	StaffId     string
	Status      string
}
