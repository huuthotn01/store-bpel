package controller

import (
	"context"
	"math/rand"
	"store-bpel/account_service/config"
	adapter2 "store-bpel/account_service/internal/adapter"
	"store-bpel/account_service/internal/repository"
	"store-bpel/account_service/schema"
	"store-bpel/library/kafka_lib"
	"time"

	"gorm.io/gorm"
)

type IAccountServiceController interface {
	GetListAccount(ctx context.Context, username string) ([]*schema.GetListAccountResponseData, error)
	AddAccount(ctx context.Context, request *schema.AddAccountRequest) error
	UpdateRole(ctx context.Context, username string, request *schema.UpdateRoleRequest) error
	SignIn(ctx context.Context, request *schema.SignInRequest) (*schema.SignInResponseData, error)
	SignUp(ctx context.Context, request *schema.SignUpRequest) error
	ChangePassword(ctx context.Context, request *schema.ChangePasswordRequest) error
}

type accountServiceController struct {
	cfg        *config.Config
	repository repository.IAccountServiceRepository

	staffAdapter    adapter2.IStaffServiceAdapter
	customerAdapter adapter2.ICustomerServiceAdapter
	kafkaAdapter    kafka_lib.IKafkaLib
}

func NewController(cfg *config.Config, db *gorm.DB) IAccountServiceController {
	// seed random math
	rand.Seed(time.Now().UnixNano())

	// init repository
	repo := repository.NewRepository(db)

	// init staff adapter
	staffAdapter := adapter2.NewStaffAdapter(cfg)

	// init customer adapter
	customerAdapter := adapter2.NewCustomerAdapter(cfg)

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
