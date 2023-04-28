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
	"store-bpel/bff/customer_bff/config"
)

type IAccountServiceAdapter interface {
	ChangePassword(ctx context.Context, request *schema.ChangePasswordRequest) error
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

func (a *accountServiceAdapter) ChangePassword(ctx context.Context, request *schema.ChangePasswordRequest) error {
	if request.Username == "" {
		err := errors.New("username must not be empty")
		log.Printf("BFF-Adapter-AccountServiceAdapter-ChangePassword error %v", err)
		return err
	}

	log.Println("Start to call account service for ChangePassword")
	defer log.Println("End call account service for ChangePassword")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/account-service/password", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-ChangePassword-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-ChangePassword-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-ChangePassword-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-ChangePassword-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-ChangePassword-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
