package controller

import (
	"context"
	"gorm.io/gorm"
	"math/rand"
	"store-bpel/account_service/adapter"
	"store-bpel/account_service/config"
	"store-bpel/account_service/repository"
	"store-bpel/account_service/schema"
	"store-bpel/library/kafka_lib"
	"time"
)

type IAccountServiceController interface {
	GetListAccount(ctx context.Context, username string) ([]*schema.GetListAccountResponseData, error)
	AddAccount(ctx context.Context, request *schema.AddAccountRequest) error
	UpdateRole(ctx context.Context, username string, request *schema.UpdateRoleRequest) error
	SignIn(ctx context.Context, request *schema.SignInRequest) (*schema.SignInResponseData, error)
	SignUp(ctx context.Context, request *schema.SignUpRequest) error
}

type accountServiceController struct {
	cfg        *config.Config
	repository repository.IAccountServiceRepository

	staffAdapter    adapter.IStaffServiceAdapter
	customerAdapter adapter.ICustomerServiceAdapter
	kafkaAdapter    kafka_lib.IKafkaLib
}

func NewController(cfg *config.Config, db *gorm.DB) IAccountServiceController {
	// seed random math
	rand.Seed(time.Now().UnixNano())

	// init repository
	repo := repository.NewRepository(db)

	// init staff adapter
	staffAdapter := adapter.NewStaffAdapter(cfg)

	// init customer adapter
	customerAdapter := adapter.NewCustomerAdapter(cfg)

	// init kafka adapter
	kafkaAdapter := kafka_lib.NewKafkaLib()

	return &accountServiceController{
		cfg:             cfg,
		repository:      repo,
		staffAdapter:    staffAdapter,
		customerAdapter: customerAdapter,
		kafkaAdapter:    kafkaAdapter,
	}
}
