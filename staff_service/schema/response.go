package schema

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
	Salary      int
	StartDate   time.Time
	EndDate     sql.NullTime
	Gender      int
	PhoneNumber string
}

type GetStaffAttendanceResponseData struct {
	AttendanceDate string
	CheckinTime    time.Time
	CheckoutTime   sql.NullTime
}
