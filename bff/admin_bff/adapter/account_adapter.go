package adapter

import (
	"net/http"
	"store-bpel/bff/admin_bff/config"
)

type IAccountServiceAdapter interface {
}

type accountServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewAccountAdapter(cfg *config.Config) IAccountServiceAdapter {
	return &accountServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.BranchServicePort,
	}
}
