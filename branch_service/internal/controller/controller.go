package controller

import (
	"context"
	"store-bpel/branch_service/config"
	repo "store-bpel/branch_service/internal/repository"
	"store-bpel/branch_service/schema"

	"gorm.io/gorm"
)

type IBranchServiceController interface {
	GetBranch(ctx context.Context) ([]*schema.GetBranchResponseData, error)
	GetBranchDetail(ctx context.Context, branchId string) (*schema.GetBranchResponseData, error)
	GetBranchStaff(ctx context.Context, branchId string) ([]string, error)
	AddBranchStaff(ctx context.Context, request *schema.AddBranchStaffRequest) error
	UpdateBranch(ctx context.Context, request *schema.UpdateBranchRequest, branchId string) error
	UpdateBranchManager(ctx context.Context, request *schema.UpdateBranchManagerRequest, branchId string) error
	AddBranch(ctx context.Context, request *schema.AddBranchRequest) error
	DeleteBranch(ctx context.Context, branchId string) error
	UploadBranchImage(ctx context.Context, branchId string) error
}

type branchServiceController struct {
	cfg        *config.Config
	repository repo.IBranchServiceRepository
}

func NewController(cfg *config.Config, db *gorm.DB) IBranchServiceController {
	// init repository
	repository := repo.NewRepository(db)

	return &branchServiceController{
		cfg:        cfg,
		repository: repository,
	}
}
