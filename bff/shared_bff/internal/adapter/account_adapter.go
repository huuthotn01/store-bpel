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
	"store-bpel/bff/shared_bff/config"
)

type IAccountServiceAdapter interface {
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
		port:       cfg.AccountServicePort,
	}
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
