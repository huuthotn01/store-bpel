package branch_service

import "time"

type GetResponse struct {
	StatusCode int         `xml:"StatusCode"`
	Message    string      `xml:"Message"`
	Data       interface{} `xml:"Data"`
}

type UpdateResponse struct {
	StatusCode int    `xml:"StatusCode"`
	Message    string `xml:"Message"`
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
