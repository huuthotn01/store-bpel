package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/goods_service/config"
	"store-bpel/order_service/schema"
)

type IOrderServiceAdapter interface {
	GetBestSellingGoods(ctx context.Context) ([]string, error)
}

type orderServiceAdapter struct {
	httpClient *http.Client
	host       string
	port       int
}

func NewOrderAdapter(cfg *config.Config) IOrderServiceAdapter {
	host := "localhost"
	if cfg.Env != "local" {
		host = "order-service"
	}
	return &orderServiceAdapter{
		host:       host,
		httpClient: &http.Client{},
		port:       cfg.OrderServicePort,
	}
}

func (a *orderServiceAdapter) GetBestSellingGoods(ctx context.Context) ([]string, error) {
	log.Printf("Start to call order service for GetBestSellingGoods")
	defer log.Println("End call order service for GetBestSellingGoods")

	var result *schema.GetBestSellingGoodsResponse

	url := fmt.Sprintf("http://%s:%d/api/order-service/best-goods", a.host, a.port)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetBestSellingGoods-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetBestSellingGoods-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetBestSellingGoods-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetBestSellingGoods-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}
