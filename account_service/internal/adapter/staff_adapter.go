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
	"store-bpel/staff_service/schema"
)

type IStaffServiceAdapter interface {
	GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error)
}

type staffServiceAdapter struct {
	httpClient *http.Client
	host       string
	port       int
}

func NewStaffAdapter(cfg *config.Config) IStaffServiceAdapter {
	host := "localhost"
	if cfg.Env != "local" {
		host = "staff-service"
	}
	return &staffServiceAdapter{
		httpClient: &http.Client{},
		host:       host,
		port:       cfg.StaffServicePort,
	}
}

func (a *staffServiceAdapter) GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error) {
	if staffId == "" {
		return nil, errors.New("[Account-StaffAdapter-GetDetailStaff] staffId must not be empty")
	}

	log.Printf("Start to call staff service for GetDetailStaff with staffId %s", staffId)
	defer log.Println("End call staff service for GetDetailStaff")
	var result *schema.GetStaffResponse
	url := fmt.Sprintf("http://%s:%d/api/staff-service/staff/%s", a.host, a.port, staffId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respByteArr, &result)
	return result.Data[0], err
}
