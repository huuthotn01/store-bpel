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
	"store-bpel/account_service/schema"
	"store-bpel/staff_service/config"
)

type IAccountServiceAdapter interface {
	CreateAccount(ctx context.Context, request *schema.AddAccountRequest) error
}

type accountServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewAccountAdapter(cfg *config.Config) IAccountServiceAdapter {
	return &accountServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.AccountServicePort,
	}
}

func (a *accountServiceAdapter) CreateAccount(ctx context.Context, request *schema.AddAccountRequest) error {
	log.Println("Start to call account service for AddAccount")
	defer log.Println("End call account service for AddAccount")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/account-service", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-AddAccount-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-AddAccount-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-AddAccount-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-AddAccount-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-AddAccount-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
