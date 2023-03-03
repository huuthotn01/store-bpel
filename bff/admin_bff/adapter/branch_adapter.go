package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/branch_service/schema"
)

type IBranchServiceAdapter interface {
	GetBranch(ctx context.Context, branchId string) (*schema.GetBranchDetailResponse, error)
}

type branchServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewBranchAdapter(cfg *config.Config) IBranchServiceAdapter {
	return &branchServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.BranchServicePort,
	}
}

func (b *branchServiceAdapter) GetBranch(ctx context.Context, branchId string) (*schema.GetBranchDetailResponse, error) {
	log.Println("Start to call branch service for GetBranch")
	defer log.Println("End call branch service for GetBranch")
	var result *schema.GetBranchDetailResponse
	url := fmt.Sprintf("http://localhost:%d/api/branch-service/%s", b.port, branchId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := b.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respByteArr, &result)
	return result, err
}
