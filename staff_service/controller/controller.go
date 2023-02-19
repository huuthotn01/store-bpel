package controller

import (
	"context"
	"store-bpel/staff_service/config"
	repo "store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"

	"gorm.io/gorm"
)

type IStaffServiceController interface {
	GetStaff(ctx context.Context, staffName, staffId string) ([]*schema.GetStaffResponseData, error)
	AddStaff(ctx context.Context, request *schema.AddStaffRequest) error
	GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error)
	GetStaffAttendance(ctx context.Context, staffId string) ([]*schema.GetStaffAttendanceResponseData, error)
	GetRequest(ctx context.Context) ([]*schema.GetRequestResponseData, error)
	CreateAddRequest(ctx context.Context, request *schema.CreateAddRequest) error
	CreateDeleteRequest(ctx context.Context, staffId string) error
	UpdateRequestStatus(ctx context.Context, request *schema.UpdateRequestStatusRequest, requestId string) error
	UpdateStaff(ctx context.Context, request *schema.UpdateStaffRequest, staffId string) error
	DeleteStaff(ctx context.Context, staffId string) error
}

type staffServiceController struct {
	db         *gorm.DB
	config     *config.Config
	repository repo.IStaffServiceRepository
}

func NewController(config *config.Config, db *gorm.DB) IStaffServiceController {
	// init repository
	repository := repo.NewRepository(db)

	return &staffServiceController{
		db:         db,
		config:     config,
		repository: repository,
	}
}
