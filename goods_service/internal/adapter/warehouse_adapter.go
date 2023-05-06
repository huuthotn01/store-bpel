package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/goods_service/config"
	"store-bpel/warehouse_service/schema"
)

type IWarehouseServiceAdapter interface {
	GetWarehouse(ctx context.Context) (*schema.UpdateResponse, error)
}

type warehouseServiceAdapter struct {
	httpClient *http.Client
	host       string
	port       int
}

func NewWarehouseAdapter(cfg *config.Config) IWarehouseServiceAdapter {
	host := "localhost"
	if cfg.Env != "local" {
		host = "warehouse-service"
	}
	return &warehouseServiceAdapter{
		httpClient: &http.Client{},
		host:       host,
		port:       cfg.WarehouseServicePort,
	}
}

func (a *warehouseServiceAdapter) GetWarehouse(ctx context.Context) (*schema.UpdateResponse, error) {
	log.Println("Start to call warehouse service for GetWarehouse")
	defer log.Println("End call warehouse service for GetWarehouse")
	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://%s:%d/api/warehouse-service/warehouse", a.host, a.port)
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
	return result, err
}
