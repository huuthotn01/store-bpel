package account_service

import (
	"context"
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/account_service"
)

type IAccountBffController interface {
	AddAccount(ctx context.Context, request *account_service.AddAccountRequest) error
	GetListAccount(ctx context.Context, request *account_service.GetListAccountRequest) ([]*account_service.GetListAccountResponseData, error)
	UpdateRole(ctx context.Context, request *account_service.UpdateRoleRequest) error
	SignIn(ctx context.Context, request *account_service.SignInRequest) (*account_service.SignInResponseData, error)
	SignUp(ctx context.Context, request *account_service.SignUpRequest) error
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
