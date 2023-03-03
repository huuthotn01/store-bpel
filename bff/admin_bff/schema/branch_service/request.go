package branch_service

type BranchRequest struct {
	Body interface{}
}

type GetBranchRequestData struct {
	BranchId string `xml:"BranchId"`
}

type AddBranchRequestData struct {
	Name     string `xml:"Name"`
	Street   string `xml:"Street"`
	Ward     string `xml:"Ward"`
	District string `xml:"District"`
	Province string `xml:"Province"`
	Open     string `xml:"Open"`
	Close    string `xml:"Close"`
}

type UpdateBranchRequestData struct {
	BranchId string `xml:"BranchId"`
	Name     string `xml:"Name"`
	Street   string `xml:"Street"`
	Ward     string `xml:"Ward"`
	District string `xml:"District"`
	Province string `xml:"Province"`
	Open     string `xml:"Open"`
	Close    string `xml:"Close"`
}
