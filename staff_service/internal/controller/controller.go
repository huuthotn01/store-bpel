package controller

import (
	"context"
	"store-bpel/library/kafka_lib"
	"store-bpel/staff_service/config"
	"store-bpel/staff_service/internal/adapter"
	repo "store-bpel/staff_service/internal/repository"
	"store-bpel/staff_service/schema"

	"gorm.io/gorm"
)

type IStaffServiceController interface {
	GetStaff(ctx context.Context, staffName, staffId string) ([]*schema.GetStaffResponseData, error)
	AddStaff(ctx context.Context, request *schema.AddStaffRequest) error
	GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error)
	GetStaffAttendance(ctx context.Context, staffId string) ([]*schema.GetStaffAttendanceResponseData, error)
	GetRequest(ctx context.Context) ([]*schema.GetRequestResponseData, error)
	CreateAddRequest(ctx context.Context, request *schema.CreateAddRequest) error
	CreateDeleteRequest(ctx context.Context, staffId string) error
	UpdateRequestStatus(ctx context.Context, request *schema.UpdateRequestStatusRequest, requestId string) error
	UpdateStaff(ctx context.Context, request *schema.UpdateStaffRequest, staffId string) error
	DeleteStaff(ctx context.Context, staffId string) error
}

type staffServiceController struct {
	db             *gorm.DB
	config         *config.Config
	accountAdapter adapter.IAccountServiceAdapter
	kafkaAdapter   kafka_lib.IKafkaLib
	repository     repo.IStaffServiceRepository
}

func NewController(config *config.Config, db *gorm.DB) IStaffServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init account adapter
	accountAdapter := adapter.NewAccountAdapter(config)

	// init kafka adapter
	kafkaAdapter := kafka_lib.NewKafkaLib()

	return &staffServiceController{
		db:             db,
		config:         config,
		accountAdapter: accountAdapter,
		kafkaAdapter:   kafkaAdapter,
		repository:     repository,
	}
}
