package account_service

import (
	"context"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/internal/adapter"
	"store-bpel/bff/shared_bff/schema/account_service"
)

type IAccountBffController interface {
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
