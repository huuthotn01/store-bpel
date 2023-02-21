package controller

import (
	"context"
	"gorm.io/gorm"
	"math/rand"
	"store-bpel/account_service/adapter"
	"store-bpel/account_service/config"
	"store-bpel/account_service/repository"
	"store-bpel/account_service/schema"
	"time"
)

type IAccountServiceController interface {
	GetListAccount(ctx context.Context, username string) ([]*schema.GetListAccountResponseData, error)
	AddAccount(ctx context.Context, request *schema.AddAccountRequest) error
	UpdateRole(ctx context.Context, username string, request *schema.UpdateRoleRequest) error
	SignIn(ctx context.Context, request *schema.SignInRequest) (int, error)
}

type accountServiceController struct {
	cfg        *config.Config
	repository repository.IAccountServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) IAccountServiceController {
	// seed random math
	rand.Seed(time.Now().UnixNano())

	// init repository
	repo := repository.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &accountServiceController{
		cfg:          cfg,
		repository:   repo,
		kafkaAdapter: kafkaAdapter,
	}
}
