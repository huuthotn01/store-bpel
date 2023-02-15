package repository

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type staffServiceRepository struct {
	db                  *gorm.DB
	accountTableName    string
	staffTableName      string
	attendanceTableName string
	requestsTaleName    string
}

type AccountModel struct {
	Username string
	StaffId  string
}

type StaffModel struct {
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
	EndDate       sql.NullTime
	Salary        int
	Gender        string
	Phone         string
}

type AttendanceModel struct {
	StaffId        string
	AttendanceDate string
	CheckinTime    time.Time
	CheckoutTime   sql.NullTime
}

type RequestsModel struct {
	Id          string
	RequestDate time.Time
	RequestType int // 1 for ADD, 2 for DELETE
	StaffId     string
	Status      string
}
