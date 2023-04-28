package account_service

import (
	"context"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/internal/adapter"
	"store-bpel/bff/customer_bff/schema/account_service"
)

type IAccountBffController interface {
	ChangePassword(ctx context.Context, request *account_service.ChangePasswordRequest) error
}

type accountBffController struct {
	cfg            *config.Config
	accountAdapter adapter.IAccountServiceAdapter
}

func NewController(cfg *config.Config) IAccountBffController {
	// init account adapter
	accountAdapter := adapter.NewAccountAdapter(cfg)

	return &accountBffController{
		cfg:            cfg,
		accountAdapter: accountAdapter,
	}
}
