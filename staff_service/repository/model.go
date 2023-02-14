package repository

import (
	"gorm.io/gorm"
	"time"
)

type staffServiceRepository struct {
	db *gorm.DB
	accountTableName string
	staffTableName string
	attendanceTableName string
	requestsTaleName string
}

type AccountModel struct {
	Username string
	StaffId string
}

type StaffModel struct {
	StaffId string
	StaffName string
	Province string
	District string
	Ward string
	Street string
	CitizenId string
	StaffPosition int
	StartDate time.Time
	Salary int
	Gender int
	Phone string
	Email string
}

type AttendanceModel struct {
	StaffId string
	AttendanceDate time.Time
	CheckinTime time.Time
	CheckoutTime time.Time
}

type RequestsModel struct {
	Id string
	RequestDate time.Time
	RequestType int
	StaffId string
	Status string
}
