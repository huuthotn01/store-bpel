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
	"store-bpel/bff/customer_bff/config"
	"store-bpel/customer_service/schema"
)

type ICustomerServiceAdapter interface {
	GetCustomer(ctx context.Context, username string) (*schema.GetCustomerInfoResponseData, error)
	AddCustomer(ctx context.Context, request *schema.AddCustomerRequest) error
	UpdateCustomer(ctx context.Context, username string, request *schema.UpdateCustomerInfoRequest) error
	UploadImage(ctx context.Context, request *schema.UploadImageRequest) error
	DeleteImage(ctx context.Context, username string) error
}

type customerServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewCustomerAdapter(cfg *config.Config) ICustomerServiceAdapter {
	return &customerServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.CustomerServicePort,
	}
}

func (a *customerServiceAdapter) GetCustomer(ctx context.Context, username string) (*schema.GetCustomerInfoResponseData, error) {
	if username == "" {
		return nil, errors.New("[BFF-Adapter-CustomerServiceAdapter-GetCustomer] username must not be empty")
	}

	log.Printf("Start to call customer service for GetCustomer, username %s", username)
	defer log.Println("End call customer service for GetCustomer")

	var result *schema.GetCustomerInfoResponse

	url := fmt.Sprintf("http://localhost:%d/api/customer-service/customer/%s", a.port, username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-GetCustomer-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-GetCustomer-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-GetCustomer-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-GetCustomer-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *customerServiceAdapter) AddCustomer(ctx context.Context, request *schema.AddCustomerRequest) error {
	log.Println("Start to call customer service for AddCustomer")
	defer log.Println("End call customer service for AddCustomer")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/customer-service/customer", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-AddCustomer-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-AddCustomer-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-AddCustomer-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-AddCustomer-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-AddCustomer-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *customerServiceAdapter) UpdateCustomer(ctx context.Context, username string, request *schema.UpdateCustomerInfoRequest) error {
	if username == "" {
		err := errors.New("[BFF-Adapter-CustomerServiceAdapter-UpdateCustomer] username must not be empty")
		log.Printf("error %v", err)
		return err
	}

	log.Println("Start to call customer service for UpdateCustomer")
	defer log.Println("End call customer service for UpdateCustomer")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/customer-service/customer/%s", a.port, username)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UpdateCustomer-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UpdateCustomer-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UpdateCustomer-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UpdateCustomer-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UpdateCustomer-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *customerServiceAdapter) UploadImage(ctx context.Context, request *schema.UploadImageRequest) error {
	log.Println("Start to call customer service for UploadImage")
	defer log.Println("End call customer service for UploadImage")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/customer-service/image", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UploadImage-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UploadImage-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UploadImage-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UploadImage-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-UploadImage-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *customerServiceAdapter) DeleteImage(ctx context.Context, username string) error {
	if username == "" {
		return errors.New("[BFF-Adapter-CustomerServiceAdapter-DeleteImage] username must not be empty")
	}

	log.Println("Start to call customer service for DeleteImage")
	defer log.Println("End call customer service for DeleteImage")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/customer-service/image/%s", a.port, username)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-DeleteImage-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-DeleteImage-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-DeleteImage-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CustomerServiceAdapter-DeleteImage-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
