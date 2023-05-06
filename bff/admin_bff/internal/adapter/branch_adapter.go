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
	GetBranchDetail(ctx context.Context, branchId string) (*schema.GetBranchResponseData, error)
	GetBranch(ctx context.Context) ([]*schema.GetBranchResponseData, error)
	AddBranch(ctx context.Context, request *schema.AddBranchRequest) error
	UpdateBranch(ctx context.Context, branchId string, request *schema.UpdateBranchRequest) error
	UpdateBranchManager(ctx context.Context, branchId string, request *schema.UpdateBranchManagerRequest) error
	DeleteBranch(ctx context.Context, branchId string) error
	GetBranchStaff(ctx context.Context, branchId string) ([]string, error)
}

type branchServiceAdapter struct {
	httpClient *http.Client
	host       string
	port       int
}

func NewBranchAdapter(cfg *config.Config) IBranchServiceAdapter {
	host := "localhost"
	if cfg.Env != "local" {
		host = "branch-service"
	}
	return &branchServiceAdapter{
		httpClient: &http.Client{},
		host:       host,
		port:       cfg.BranchServicePort,
	}
}

func (b *branchServiceAdapter) GetBranchDetail(ctx context.Context, branchId string) (*schema.GetBranchResponseData, error) {
	if branchId == "" {
		err := errors.New("branchId must not be empty")
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchDetail error %v", err)
		return nil, err
	}

	log.Printf("Start to call branch service for GetBranchDetail, branchId %s", branchId)
	defer log.Println("End call branch service for GetBranchDetail")

	var (
		result = &schema.GetBranchDetailResponse{}
	)

	url := fmt.Sprintf("http://%s:%d/api/branch-service/%s", b.host, b.port, branchId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchDetail-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := b.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchDetail-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchDetail-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchDetail-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (b *branchServiceAdapter) GetBranch(ctx context.Context) ([]*schema.GetBranchResponseData, error) {
	log.Printf("Start to call branch service for GetBranch")
	defer log.Println("End call branch service for GetBranch")

	var (
		result = &schema.GetBranchResponse{}
	)

	url := fmt.Sprintf("http://%s:%d/api/branch-service", b.host, b.port)
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

	return result.Data, nil
}

func (b *branchServiceAdapter) AddBranch(ctx context.Context, request *schema.AddBranchRequest) error {
	log.Println("Start to call branch service for AddBranch")
	defer log.Println("End call branch service for AddBranch")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/branch-service", b.host, b.port)
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

func (b *branchServiceAdapter) UpdateBranch(ctx context.Context, branchId string, request *schema.UpdateBranchRequest) error {
	log.Println("Start to call branch service for UpdateBranch")
	defer log.Println("End call branch service for UpdateBranch")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/branch-service/%s", b.host, b.port, branchId)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranch-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranch-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := b.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranch-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranch-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranch-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (b *branchServiceAdapter) UpdateBranchManager(ctx context.Context, branchId string, request *schema.UpdateBranchManagerRequest) error {
	log.Println("Start to call branch service for UpdateBranchManager")
	defer log.Println("End call branch service for UpdateBranchManager")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/branch-service/manager/%s", b.host, b.port, branchId)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranchManager-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranchManager-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := b.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranchManager-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranchManager-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-UpdateBranchManager-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (b *branchServiceAdapter) DeleteBranch(ctx context.Context, branchId string) error {
	log.Println("Start to call branch service for DeleteBranch")
	defer log.Println("End call branch service for DeleteBranch")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/branch-service/%s", b.host, b.port, branchId)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-DeleteBranch-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := b.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-DeleteBranch-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-DeleteBranch-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-DeleteBranch-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (b *branchServiceAdapter) GetBranchStaff(ctx context.Context, branchId string) ([]string, error) {
	if branchId == "" {
		err := errors.New("branchId must not be empty")
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchStaff error %v", err)
		return nil, err
	}

	log.Printf("Start to call branch service for GetBranchStaff, branchId %s", branchId)
	defer log.Println("End call branch service for GetBranchStaff")

	var (
		result = &schema.GetBranchStaffResponse{}
	)

	url := fmt.Sprintf("http://%s:%d/api/branch-service/staff/%s", b.host, b.port, branchId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchStaff-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := b.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchStaff-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchStaff-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-BranchServiceAdapter-GetBranchStaff-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}
