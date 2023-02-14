package controller

import (
	"gorm.io/gorm"
	"store-bpel/customer_invoice_service/adapter"
	"store-bpel/customer_invoice_service/config"
	repo "store-bpel/customer_invoice_service/repository"
	// "store-bpel/customer_invoice_service/schema"
)

type ICustomerInvoiceServiceController interface {
}

type customerInvoiceServiceController struct {
	cfg        *config.Config
	repository repo.ICustomerInvoiceServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) ICustomerInvoiceServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &customerInvoiceServiceController{
		cfg:          cfg,
		repository:   repository,
		kafkaAdapter: kafkaAdapter,
	}
}
