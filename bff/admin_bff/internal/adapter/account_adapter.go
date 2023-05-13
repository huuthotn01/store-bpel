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
	"store-bpel/bff/admin_bff/config"
)

type IAccountServiceAdapter interface {
	GetListAccount(ctx context.Context, username string) ([]*schema.GetListAccountResponseData, error)
	AddAccount(ctx context.Context, request *schema.AddAccountRequest) error
	UpdateRole(ctx context.Context, username string, request *schema.UpdateRoleRequest) error
	ChangePassword(ctx context.Context, request *schema.ChangePasswordRequest) error
}

type accountServiceAdapter struct {
	httpClient *http.Client
	host       string
	port       int
}

func NewAccountAdapter(cfg *config.Config) IAccountServiceAdapter {
	return &accountServiceAdapter{
		httpClient: &http.Client{},
		host:       cfg.AccountServiceHost,
		port:       cfg.AccountServicePort,
	}
}

func (a *accountServiceAdapter) GetListAccount(ctx context.Context, username string) ([]*schema.GetListAccountResponseData, error) {
	log.Printf("Start to call account service for GetListAccount, filtering username %s", username)
	defer log.Println("End call account service for GetListAccount")

	var result *schema.GetListAccountResponse

	url := fmt.Sprintf("http://%s:%d/api/account-service", a.host, a.port)
	if username != "" {
		url += fmt.Sprintf("?username=%s", username)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-GetListAccount-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-GetListAccount-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-GetListAccount-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-GetListAccount-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *accountServiceAdapter) AddAccount(ctx context.Context, request *schema.AddAccountRequest) error {
	log.Println("Start to call account service for AddAccount")
	defer log.Println("End call account service for AddAccount")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/account-service", a.host, a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-AddAccount-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-AddAccount-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-AddAccount-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-AddAccount-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-AddAccount-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *accountServiceAdapter) UpdateRole(ctx context.Context, username string, request *schema.UpdateRoleRequest) error {
	if username == "" {
		err := errors.New("username must not be empty")
		log.Printf("BFF-Adapter-AccountServiceAdapter-UpdateRole error %v", err)
		return err
	}

	log.Println("Start to call account service for UpdateRole")
	defer log.Println("End call account service for UpdateRole")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/account-service/%s", a.host, a.port, username)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-UpdateRole-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-UpdateRole-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-UpdateRole-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-UpdateRole-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-UpdateRole-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
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
	url := fmt.Sprintf("http://%s:%d/api/account-service/password", a.host, a.port)
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
