package schema

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

type UpdateStaffRequest struct {
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
	Status string
}
