package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/staff_service/config"
	repo "store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"
)

type IStaffServiceController interface {
	GetStaff(ctx context.Context) ([]*schema.GetStaffResponseData, error)
	AddStaff(ctx context.Context, request *schema.AddStaffRequest) error
	GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error)
	GetStaffAttendance(ctx context.Context, staffId string) ([]*schema.GetStaffAttendanceResponseData, error)
	CreateAddRequest(ctx context.Context, request *schema.CreateAddRequest) error
	DeleteAddRequest(ctx context.Context, staffId string) error
	UpdateRequestStatus(ctx context.Context, request *schema.UpdateRequestStatusRequest, requestId string) error
}

type staffServiceController struct {
	config     *config.Config
	repository repo.IStaffServiceRepository
}

func NewController(config *config.Config, db *gorm.DB) IStaffServiceController {
	// init repository
	repository := repo.NewRepository(db)

	return &staffServiceController{
		config:     config,
		repository: repository,
	}
}
