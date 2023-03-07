package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/branch_service/adapter"
	"store-bpel/branch_service/config"
	repo "store-bpel/branch_service/repository"
	"store-bpel/branch_service/schema"
)

type IBranchServiceController interface {
	GetBranch(ctx context.Context) ([]*schema.GetBranchResponseData, error)
	GetBranchDetail(ctx context.Context, branchId string) (*schema.GetBranchResponseData, error)
	GetBranchStaff(ctx context.Context, branchId string) ([]string, error)
	UpdateBranch(ctx context.Context, request *schema.UpdateBranchRequest, branchId int32) error
	UpdateBranchManager(ctx context.Context, request *schema.UpdateBranchManagerRequest, branchId int32) error
	AddBranch(ctx context.Context, request *schema.AddBranchRequest) error
	DeleteBranch(ctx context.Context, branchId int32) error
	UploadBranchImage(ctx context.Context, branchId string) error
}

type branchServiceController struct {
	cfg        *config.Config
	repository repo.IBranchServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) IBranchServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &branchServiceController{
		cfg:          cfg,
		repository:   repository,
		kafkaAdapter: kafkaAdapter,
	}
}
