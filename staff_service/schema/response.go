package schema

import (
	"database/sql"
	"time"
)

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetStaffResponse struct {
	StatusCode int
	Message    string
	Data       []*GetStaffResponseData
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

type GetStaffAttendanceResponse struct {
	StatusCode int
	Message    string
	Data       []*GetStaffAttendanceResponseData
}

type GetStaffAttendanceResponseData struct {
	AttendanceDate string
	CheckinTime    time.Time
	CheckoutTime   sql.NullTime
}

type GetRequestResponse struct {
	StatusCode int
	Message    string
	Data       []*GetRequestResponseData
}

type GetRequestResponseData struct {
	Id            string
	RequestDate   time.Time
	RequestType   string // ADD or DELETE
	Status        string
	StaffId       string
	StaffName     string
	Province      string
	District      string
	Ward          string
	Street        string
	Hometown      string
	CitizenId     string
	StaffPosition string
	Birthdate     string
	StartDate     time.Time
	Salary        int
	Gender        string
	Phone         string
	Email         string
	BranchId      string
}
