package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/goods_service/schema"
	"store-bpel/order_service/config"
)

type IGoodsServiceAdapter interface {
	GetProductDetail(ctx context.Context, productId string) (*schema.GetGoodsDefaultResponseData, error)
}

type goodsServiceAdapter struct {
	httpClient *http.Client
	host       string
	port       int
}

func NewGoodsAdapter(cfg *config.Config) IGoodsServiceAdapter {
	return &goodsServiceAdapter{
		httpClient: &http.Client{},
		host:       cfg.GoodsServiceHost,
		port:       cfg.GoodsServicePort,
	}
}

func (a *goodsServiceAdapter) GetProductDetail(ctx context.Context, productId string) (*schema.GetGoodsDefaultResponseData, error) {
	if productId == "" {
		return nil, errors.New("[BFF-Adapter-GoodsServiceAdapter-GetProductDetail] productId must not be empty")
	}

	log.Println("Start to call goods service for GetProductDetail")
	defer log.Println("End call goods service for GetProductDetail")

	var result *schema.GetDetailProductsResponse
	url := fmt.Sprintf("http://%s:%d/api/goods-service/product/%s", a.host, a.port, productId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetProductDetail-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetProductDetail-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetProductDetail-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetProductDetail-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}
