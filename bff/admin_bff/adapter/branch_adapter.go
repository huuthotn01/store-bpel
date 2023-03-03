package adapter

import (
	"context"
	"encoding/json"
	"errors"
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
	if branchId == "" {
		err := errors.New("branchId must not be empty")
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranch error %v", err)
		return nil, err
	}
	log.Printf("Start to call branch service for GetBranch, branchId %s", branchId)
	defer log.Println("End call branch service for GetBranch")
	var result *schema.GetBranchDetailResponse
	url := fmt.Sprintf("http://localhost:%d/api/branch-service/%s", b.port, branchId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranch-NewRequestWithContext error %v", err)
		return nil, err
	}
	resp, err := b.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranch-httpClient.Do error %v", err)
		return nil, err
	}
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranch-ioutil.ReadAll error %v", err)
		return nil, err
	}
	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranch-json.Unmarshal error %v", err)
		return nil, err
	}
	return result, nil
}
