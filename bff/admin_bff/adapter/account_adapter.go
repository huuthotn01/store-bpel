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
	SignIn(ctx context.Context, request *schema.SignInRequest) (*schema.SignInResponseData, error)
	SignUp(ctx context.Context, request *schema.SignUpRequest) error
}

type accountServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewAccountAdapter(cfg *config.Config) IAccountServiceAdapter {
	return &accountServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.BranchServicePort,
	}
}

func (a *accountServiceAdapter) GetListAccount(ctx context.Context, username string) ([]*schema.GetListAccountResponseData, error) {
	log.Printf("Start to call account service for GetListAccount, filtering username %s", username)
	defer log.Println("End call account service for GetListAccount")

	var result *schema.GetListAccountResponse

	url := fmt.Sprintf("http://localhost:%d/api/account-service", a.port)
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
	url := fmt.Sprintf("http://localhost:%d/api/account-service", a.port)
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
	url := fmt.Sprintf("http://localhost:%d/api/account-service/%s", a.port, username)
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

func (a *accountServiceAdapter) SignIn(ctx context.Context, request *schema.SignInRequest) (*schema.SignInResponseData, error) {
	log.Println("Start to call account service for SignIn")
	defer log.Println("End call account service for SignIn")

	var result *schema.SignInResponse
	url := fmt.Sprintf("http://localhost:%d/api/account-service/sign-in", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignIn-Marshal error %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignIn-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignIn-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignIn-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignIn-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *accountServiceAdapter) SignUp(ctx context.Context, request *schema.SignUpRequest) error {
	log.Println("Start to call account service for SignUp")
	defer log.Println("End call account service for SignUp")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/account-service/sign-up", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignUp-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignUp-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignUp-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignUp-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-AccountServiceAdapter-SignUp-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
