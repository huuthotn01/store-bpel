package branch_service

import (
	"context"
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
	branch_schema "store-bpel/bff/admin_bff/schema/branch_service"
)

type IBranchBffController interface {
	GetBranch(ctx context.Context) ([]*branch_schema.GetBranchResponseData, error)
	GetBranchDetail(ctx context.Context, branchId string) (*branch_schema.GetBranchResponseData, error)
	AddBranch(ctx context.Context, request *branch_schema.AddBranchRequest) error
	UpdateBranch(ctx context.Context, request *branch_schema.UpdateBranchRequest) error
	UpdateBranchManager(ctx context.Context, request *branch_schema.UpdateBranchManagerRequest) error
	DeleteBranch(ctx context.Context, request *branch_schema.DeleteBranchRequest) error
	GetBranchStaff(ctx context.Context, request *branch_schema.GetBranchStaffRequest) (*branch_schema.GetBranchStaffResponseData, error)
}

type branchBffController struct {
	cfg           *config.Config
	branchAdapter adapter.IBranchServiceAdapter
}

func NewController(cfg *config.Config) IBranchBffController {
	// init branch adapter
	branchAdapter := adapter.NewBranchAdapter(cfg)

	return &branchBffController{
		cfg:           cfg,
		branchAdapter: branchAdapter,
	}
}
