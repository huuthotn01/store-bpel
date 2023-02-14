package schema

import "time"

type AddStaffRequest struct {
	Name         string
	Birthdate    time.Time
	Hometown     string
	CitizenId    string
	Phone        string
	Street       string
	Ward         string
	District     string
	Province     string
	WorkingPlace string // TODO post check with API docs
	Role         string // TODO post check with API docs
	Salary       int
	Username     string
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
	Salary       int
}

type UpdateRequestStatusRequest struct {
	Status string
}
