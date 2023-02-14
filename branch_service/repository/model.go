package repository

import (
	"gorm.io/gorm"
	"time"
)

type branchServiceRepository struct {
	db                     *gorm.DB
	branchTableName        string
	branchImgTableName     string
	branchManagerTableName string
	branchStaffTableName   string
}

type BranchModel struct {
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

type BranchImgModel struct {
	BranchCode int32
	BranchImg  string
}

type BranchManagerModel struct {
	BranchCode  int32
	ManagerCode string
	StartDate   time.Time
	EndDate     time.Time
}

type BranchStaffModel struct {
	BranchCode int32
	StaffCode  string
	StartDate  time.Time
	EndDate    time.Time
}
