package branch_service

type GetBranchDetailRequest struct {
	BranchId string `xml:"BranchId"`
}

type AddBranchRequest struct {
	Name     string `xml:"Name"`
	Street   string `xml:"Street"`
	Ward     string `xml:"Ward"`
	District string `xml:"District"`
	Province string `xml:"Province"`
	Open     string `xml:"Open"`
	Close    string `xml:"Close"`
}

type UpdateBranchRequest struct {
	BranchId string `xml:"BranchId"`
	Name     string `xml:"Name"`
	Street   string `xml:"Street"`
	Ward     string `xml:"Ward"`
	District string `xml:"District"`
	Province string `xml:"Province"`
	Open     string `xml:"Open"`
	Close    string `xml:"Close"`
}

type UpdateBranchManagerRequest struct {
	BranchId string `xml:"BranchId"`
	StaffId  string `xml:"StaffId"`
}

type DeleteBranchRequest struct {
	BranchId string `xml:"BranchId"`
}

type GetBranchStaffRequest struct {
	BranchId string `xml:"BranchId"`
}

type GetBranchStatisticRequest struct {
	BranchId string `xml:"BranchId"`
	FromTime string `xml:"FromTime"`
	ToTime   string `xml:"ToTime"`
}
