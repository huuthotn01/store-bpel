package controller

import (
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
)

type IAdminBffController interface {
}

type adminBffController struct {
	cfg           *config.Config
	BranchAdapter adapter.IBranchServiceAdapter
}

func NewController(cfg *config.Config) IAdminBffController {
	// init branch adapter
	branchAdapter := adapter.NewBranchAdapter(cfg)

	return &adminBffController{
		cfg:           cfg,
		BranchAdapter: branchAdapter,
	}
}
