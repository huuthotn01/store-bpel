package schema

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

type GetBranchStaffResponseData struct {
	Staffs []string
}
