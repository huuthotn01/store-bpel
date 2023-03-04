package branch_service

import (
	"context"
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
	branch_schema "store-bpel/bff/admin_bff/schema/branch_service"
	"store-bpel/branch_service/schema"
)

type IBranchBffController interface {
	GetBranch(ctx context.Context, branchId string) (*schema.GetBranchResponseData, error)
	AddBranch(ctx context.Context, request *branch_schema.AddBranchRequest) error
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
