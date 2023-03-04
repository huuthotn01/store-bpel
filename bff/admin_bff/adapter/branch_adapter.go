package adapter

import (
	"bytes"
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
	GetBranch(ctx context.Context, branchId string) (*schema.GetBranchResponseData, error)
	AddBranch(ctx context.Context, request *schema.AddBranchRequest) error
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

func (b *branchServiceAdapter) GetBranch(ctx context.Context, branchId string) (*schema.GetBranchResponseData, error) {
	if branchId == "" {
		err := errors.New("branchId must not be empty")
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranch error %v", err)
		return nil, err
	}

	log.Printf("Start to call branch service for GetBranch, branchId %s", branchId)
	defer log.Println("End call branch service for GetBranch")

	var (
		result     = &schema.GetResponse{}
		resultData = &schema.GetBranchResponseData{}
	)
	result.Data = resultData

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

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return resultData, nil
}

func (b *branchServiceAdapter) AddBranch(ctx context.Context, request *schema.AddBranchRequest) error {
	log.Println("Start to call branch service for AddBranch")
	defer log.Println("End call branch service for AddBranch")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/branch-service", b.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-AddBranch-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-AddBranch-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := b.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-AddBranch-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-AddBranch-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-AddBranch-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
