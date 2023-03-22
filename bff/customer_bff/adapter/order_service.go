package adapter

import (
	"net/http"
	"store-bpel/bff/customer_bff/config"
)

type IOrderServiceAdapter interface {
}

type orderServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewOrderAdapter(cfg *config.Config) IOrderServiceAdapter {
	return &orderServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.CustomerServicePort,
	}
}
