package staff_service

type GetStaffRequest struct {
	StaffName string
	StaffId   string
}

type AddStaffRequest struct {
	Name         string
	Birthdate    string
	Hometown     string
	CitizenId    string
	Phone        string
	Street       string
	Ward         string
	District     string
	Province     string
	WorkingPlace string
	Role         string
	Gender       string
	Salary       int
	Email        string
}

type GetStaffAttendanceRequest struct {
	StaffId string
}

type CreateDeleteRequest struct {
	StaffId string
}

type UpdateStaffRequest struct {
	StaffId      string
	Name         string
	Birthdate    string
	Hometown     string
	CitizenId    string
	Phone        string
	Province     string
	District     string
	Ward         string
	Street       string
	WorkingPlace string
	Role         string
	Gender       string
	Salary       int
}

type CreateAddRequest struct {
	Name         string
	Birthdate    string
	Hometown     string
	CitizenId    string
	Phone        string
	Province     string
	District     string
	Ward         string
	Street       string
	WorkingPlace string
	Role         string
	Gender       string
	Salary       int
	Email        string
}

type UpdateRequestStatusRequest struct {
	RequestId string
	Status    string
}
