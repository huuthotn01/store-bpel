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
	"store-bpel/warehouse_service/schema"
)

type IWarehouseServiceAdapter interface {
	GetManager(ctx context.Context, warehouseId string) (*schema.GetWarehouseManagerResponseData, error)
	GetStaff(ctx context.Context, warehouseId string) ([]*schema.GetWarehouseStaffResponseData, error)
	GetWarehouse(ctx context.Context, warehouseId string) (*schema.GetWarehouseResponseData, error)
	AddStaff(ctx context.Context, request *schema.AddWarehouseStaffRequest) error
	AddWarehouse(ctx context.Context, request *schema.AddWarehouseRequest) error
	UpdateStaff(ctx context.Context, request *schema.UpdateWarehouseStaffRequest) error
	UpdateManager(ctx context.Context, request *schema.UpdateManagerRequest) error
	UpdateWarehouse(ctx context.Context, request *schema.UpdateWarehouseRequest) error
	DeleteStaff(ctx context.Context, request *schema.DeleteWarehouseStaffRequest) error
	DeleteWarehouse(ctx context.Context, request *schema.DeleteWarehouseRequest) error
	GetAllWarehouse(ctx context.Context) ([]*schema.GetWarehouseResponseData, error)
}

type warehouseServiceAdapter struct {
	httpClient *http.Client
	host       string
	port       int
}

func NewWarehouseAdapter(cfg *config.Config) IWarehouseServiceAdapter {
	return &warehouseServiceAdapter{
		httpClient: &http.Client{},
		host:       cfg.WarehouseServiceHost,
		port:       cfg.WarehouseServicePort,
	}
}

func (a *warehouseServiceAdapter) GetManager(ctx context.Context, warehouseId string) (*schema.GetWarehouseManagerResponseData, error) {
	if warehouseId == "" {
		return nil, errors.New("[BFF-Adapter-WarehouseServiceAdapter-GetManager] warehouseId must not be empty")
	}

	log.Printf("Start to call warehouse service for GetManager, warehouseId %s", warehouseId)
	defer log.Println("End call warehouse service for GetManager")

	var result *schema.GetWarehouseManagerResponse

	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/manager/%s", a.host, a.port, warehouseId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetManager-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetManager-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetManager-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetManager-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *warehouseServiceAdapter) GetStaff(ctx context.Context, warehouseId string) ([]*schema.GetWarehouseStaffResponseData, error) {
	if warehouseId == "" {
		return nil, errors.New("[BFF-Adapter-WarehouseServiceAdapter-GetManager] warehouseId must not be empty")
	}

	log.Printf("Start to call warehouse service for GetStaff, warehouseId %s", warehouseId)
	defer log.Println("End call warehouse service for GetStaff")

	var result *schema.GetWarehouseStaffResponse

	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/staff/%s", a.host, a.port, warehouseId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetStaff-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetStaff-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetStaff-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetStaff-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *warehouseServiceAdapter) GetWarehouse(ctx context.Context, warehouseId string) (*schema.GetWarehouseResponseData, error) {
	if warehouseId == "" {
		return nil, errors.New("[BFF-Adapter-WarehouseServiceAdapter-GetManager] warehouseId must not be empty")
	}

	log.Printf("Start to call warehouse service for GetWarehouse, warehouseId %s", warehouseId)
	defer log.Println("End call warehouse service for GetWarehouse")

	var result *schema.GetWarehouseResponse

	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/warehouse/%s", a.host, a.port, warehouseId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetWarehouse-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetWarehouse-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetWarehouse-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetWarehouse-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *warehouseServiceAdapter) GetAllWarehouse(ctx context.Context) ([]*schema.GetWarehouseResponseData, error) {

	log.Printf("Start to call warehouse service for GetAllWarehouse, warehouseId")
	defer log.Println("End call warehouse service for GetAllWarehouse")

	var result *schema.GetAllWarehouseResponse

	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/warehouse", a.host, a.port)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetAllWarehouse-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetAllWarehouse-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetAllWarehouse-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-GetAllWarehouse-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *warehouseServiceAdapter) AddStaff(ctx context.Context, request *schema.AddWarehouseStaffRequest) error {
	log.Println("Start to call warehouse service for AddStaff")
	defer log.Println("End call warehouse service for AddStaff")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/staff", a.host, a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddStaff-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddStaff-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddStaff-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddStaff-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddStaff-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *warehouseServiceAdapter) AddWarehouse(ctx context.Context, request *schema.AddWarehouseRequest) error {
	log.Println("Start to call warehouse service for AddWarehouse")
	defer log.Println("End call warehouse service for AddWarehouse")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/warehouse", a.host, a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddWarehouse-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddWarehouse-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddWarehouse-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddWarehouse-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-AddWarehouse-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *warehouseServiceAdapter) UpdateStaff(ctx context.Context, request *schema.UpdateWarehouseStaffRequest) error {
	log.Println("Start to call warehouse service for UpdateStaff")
	defer log.Println("End call warehouse service for UpdateStaff")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/staff", a.host, a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateStaff-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateStaff-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateStaff-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateStaff-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateStaff-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *warehouseServiceAdapter) UpdateManager(ctx context.Context, request *schema.UpdateManagerRequest) error {
	log.Println("Start to call warehouse service for UpdateManager")
	defer log.Println("End call warehouse service for UpdateManager")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/manager", a.host, a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateManager-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateManager-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateManager-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateManager-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateManager-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *warehouseServiceAdapter) UpdateWarehouse(ctx context.Context, request *schema.UpdateWarehouseRequest) error {
	log.Println("Start to call warehouse service for UpdateWarehouse")
	defer log.Println("End call warehouse service for UpdateWarehouse")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/warehouse", a.host, a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateWarehouse-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateWarehouse-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateWarehouse-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateWarehouse-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-UpdateWarehouse-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *warehouseServiceAdapter) DeleteStaff(ctx context.Context, request *schema.DeleteWarehouseStaffRequest) error {
	log.Println("Start to call warehouse service for DeleteStaff")
	defer log.Println("End call warehouse service for DeleteStaff")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/staff", a.host, a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteStaff-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteStaff-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteStaff-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteStaff-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteStaff-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *warehouseServiceAdapter) DeleteWarehouse(ctx context.Context, request *schema.DeleteWarehouseRequest) error {
	log.Println("Start to call warehouse service for DeleteWarehouse")
	defer log.Println("End call warehouse service for DeleteWarehouse")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/warehouse", a.host, a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteWarehouse-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteWarehouse-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteWarehouse-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteWarehouse-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-WarehouseServiceAdapter-DeleteWarehouse-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
