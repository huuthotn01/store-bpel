package adapter

import (
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
	GetCustomer(ctx context.Context, username string) (*schema.GetCustomerInfoResponseData, error)
}

type customerServiceAdapter struct {
	httpClient *http.Client
	host       string
	port       int
}

func NewCustomerAdapter(cfg *config.Config) ICustomerServiceAdapter {
	return &customerServiceAdapter{
		httpClient: &http.Client{},
		host:       cfg.CustomerServiceHost,
		port:       cfg.CustomerServicePort,
	}
}

func (a *customerServiceAdapter) GetCustomer(ctx context.Context, username string) (*schema.GetCustomerInfoResponseData, error) {
	if username == "" {
		return nil, errors.New("[AccountService-CustomerServiceAdapter-GetCustomer] username must not be empty")
	}

	log.Printf("Start to call customer service for GetCustomer, username %s", username)
	defer log.Println("End call customer service for GetCustomer")

	var result *schema.GetCustomerInfoResponse

	url := fmt.Sprintf("http://%s:%d/api/customer-service/customer/%s", a.host, a.port, username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("AccountService-CustomerServiceAdapter-GetCustomer-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("AccountService-CustomerServiceAdapter-GetCustomer-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("AccountService-CustomerServiceAdapter-GetCustomer-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("AccountService-CustomerServiceAdapter-GetCustomer-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}
