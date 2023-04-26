package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/staff_service/schema"
	"store-bpel/warehouse_service/config"
)

type IStaffServiceAdapter interface {
	GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error)
}

type staffServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewStaffAdapter(cfg *config.Config) IStaffServiceAdapter {
	return &staffServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.StaffServicePort,
	}
}

func (a *staffServiceAdapter) GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error) {
	if staffId == "" {
		return nil, errors.New("[WarehouseServiceAdapter-StaffAdapter-GetDetailStaff] staffId must not be empty")
	}
	log.Println("Start to call staff service for GetStaff")
	defer log.Println("End call staff service for GetStaff")
	var result *schema.GetStaffResponse
	url := fmt.Sprintf("http://localhost:%d/api/staff-service/staff/%s", a.port, staffId)
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
