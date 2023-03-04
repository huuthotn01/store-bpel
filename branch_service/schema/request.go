package schema

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
