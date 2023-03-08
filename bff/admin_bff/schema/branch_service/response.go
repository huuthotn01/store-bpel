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
	BranchCode     int32     `xml:"BranchId"`
	BranchName     string    `xml:"Name"`
	BranchProvince string    `xml:"Province"`
	BranchDistrict string    `xml:"District"`
	BranchWard     string    `xml:"Ward"`
	BranchStreet   string    `xml:"Street"`
	CreatedAt      time.Time `xml:"CreatedAt"`
	Manager        string    `xml:"Manager"`
	Open           string    `xml:"Open"`
	Close          string    `xml:"Close"`
}

type GetBranchStaffResponseData struct {
	Staffs []string
}
