package schema

import "time"

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetWarehouseStaffResponse struct {
	StatusCode int
	Message    string
	Data       []*GetWarehouseStaffResponseData
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

type GetWarehouseManagerResponse struct {
	StatusCode int
	Message    string
	Data       *GetWarehouseManagerResponseData
}

type GetAllWarehouseManagerResponse struct {
	StatusCode int
	Message    string
	Data       []*GetWarehouseResponseData
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

type GetWarehouseResponse struct {
	StatusCode int
	Message    string
	Data       *GetWarehouseResponseData
}

type GetAllWarehouseResponse struct {
	StatusCode int
	Message    string
	Data       []*GetWarehouseResponseData
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
