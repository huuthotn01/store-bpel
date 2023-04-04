package warehouse_service

import "time"

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type GetWarehouseManagerResponseData struct {
	StaffId     string
	StaffName   string
	Street      string
	Ward        string
	District    string
	Province    string
	CitizenId   string
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

type GetWarehouseStaffResponseData struct {
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

type GetWarehouseResponseData struct {
	WarehouseCode string
	WarehouseName string
	Capacity      int
	CreatedAt     time.Time
	Street        string
	Ward          string
	District      string
	Province      string
}
