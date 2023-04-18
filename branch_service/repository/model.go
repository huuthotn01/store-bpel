package repository

import (
	"time"

	"gorm.io/gorm"
)

type branchServiceRepository struct {
	db                     *gorm.DB
	branchTableName        string
	branchImgTableName     string
	branchManagerTableName string
	branchStaffTableName   string
}

type BranchModel struct {
	BranchCode     string
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

type BranchImgModel struct {
	BranchCode string
	BranchImg  string
}

type BranchManagerModel struct {
	BranchCode  string
	ManagerCode string
	StartDate   time.Time
	EndDate     time.Time
}

type BranchStaffModel struct {
	BranchCode string
	StaffCode  string
	StartDate  time.Time
	EndDate    time.Time
}
