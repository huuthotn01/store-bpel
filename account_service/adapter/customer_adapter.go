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
	"store-bpel/account_service/config"
	"store-bpel/customer_service/schema"
)

type ICustomerServiceAdapter interface {
	AddCustomer(ctx context.Context, request *schema.AddCustomerRequest) error
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

func (a *customerServiceAdapter) AddCustomer(ctx context.Context, request *schema.AddCustomerRequest) error {
	log.Println("Start to call customer service for AddCustomer")
	defer log.Println("End call customer service for AddCustomer")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/customer-service/customer", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("Customer Service-CustomerServiceAdapter-AddCustomer-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("Customer Service-CustomerServiceAdapter-AddCustomer-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("Customer Service-CustomerServiceAdapter-AddCustomer-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Customer Service-CustomerServiceAdapter-AddCustomer-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("Customer Service-CustomerServiceAdapter-AddCustomer-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
