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
	UpdateRole(ctx context.Context, username string, request *schema.UpdateRoleRequest) error
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

func (a *accountServiceAdapter) UpdateRole(ctx context.Context, username string, request *schema.UpdateRoleRequest) error {
	log.Println("Start to call account service for UpdateRole")
	defer log.Println("End call account service for UpdateRole")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/account-service/role/%s", a.host, a.port, username)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-UpdateRole-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-UpdateRole-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-UpdateRole-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-UpdateRole-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("Staff Service-AccountServiceAdapter-UpdateRole-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
