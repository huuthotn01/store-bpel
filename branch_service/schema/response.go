package schema

import "time"

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
