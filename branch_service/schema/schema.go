package schema

import "time"

type AddBranchRequest struct {
	Name     string
	Street   string
	Ward     string
	District string
	Province string
	Open     string
	Close    string
}

type UpdateBranchRequest struct {
	Name     string
	Street   string
	Ward     string
	District string
	Province string
	Open     string
	Close    string
}

type UpdateBranchManagerRequest struct {
	StaffId string
}

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetBranchResponse struct {
	StatusCode int
	Message    string
	Data       []*GetBranchResponseData
}

type GetBranchDetailResponse struct {
	StatusCode int
	Message    string
	Data       *GetBranchResponseData
}

type GetBranchResponseData struct {
	BranchCode     int32
	BranchName     string
	BranchProvince string
	BranchDistrict string
	BranchWard     string
	BranchStreet   string
	CreatedAt      time.Time
	Manager        string
	OpenTime       string
	CloseTime      string
}

type GetBranchStaffResponse struct {
	StatusCode int
	Message    string
	Data       []string
}
