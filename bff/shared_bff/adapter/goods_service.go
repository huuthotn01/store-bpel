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
	"store-bpel/bff/shared_bff/config"
	"store-bpel/goods_service/schema"
)

type IGoodsServiceAdapter interface {
	GetGoods(ctx context.Context) ([]*schema.GetGoodsResponseData, error)
	GetGoodsDefault(ctx context.Context, request *schema.GetGoodsDefaultRequest) ([]*schema.GetGoodsDefaultResponseData, error)
	GetProductDetail(ctx context.Context, productId string) (*schema.GetGoodsDefaultResponseData, error)
	GetGoodsDetail(ctx context.Context, goodsId string) ([]*schema.GetGoodsResponseData, error)
	CheckWarehouse(ctx context.Context, request *schema.CheckWarehouseRequest) (*schema.CheckWarehouseResponseData, error)
	CreateWHTransfer(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error
}

type goodsServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewGoodsAdapter(cfg *config.Config) IGoodsServiceAdapter {
	return &goodsServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.GoodsServicePort,
	}
}

func (a *goodsServiceAdapter) GetGoods(ctx context.Context) ([]*schema.GetGoodsResponseData, error) {
	log.Println("Start to call goods service for GetGoods")
	defer log.Println("End call goods service for GetGoods")

	var result *schema.GetGoodsResponse
	url := fmt.Sprintf("http://localhost:%d/api/goods-service/goods", a.port)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoods-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoods-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoods-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoods-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *goodsServiceAdapter) GetGoodsDefault(ctx context.Context, request *schema.GetGoodsDefaultRequest) ([]*schema.GetGoodsDefaultResponseData, error) {
	log.Println("Start to call goods service for GetGoodsDefault")
	defer log.Println("End call goods service for GetGoodsDefault")

	var result *schema.GetGoodsDefaultResponse
	url := fmt.Sprintf("http://localhost:%d/api/goods-service/default-goods", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDefault-Marshal error %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDefault-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDefault-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDefault-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDefault-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *goodsServiceAdapter) GetProductDetail(ctx context.Context, productId string) (*schema.GetGoodsDefaultResponseData, error) {
	if productId == "" {
		return nil, errors.New("[BFF-Adapter-GoodsServiceAdapter-GetProductDetail] productId must not be empty")
	}

	log.Println("Start to call goods service for GetProductDetail")
	defer log.Println("End call goods service for GetProductDetail")

	var result *schema.GetDetailProductsResponse
	url := fmt.Sprintf("http://localhost:%d/api/goods-service/product/%s", a.port, productId)

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

func (a *goodsServiceAdapter) GetGoodsDetail(ctx context.Context, goodsId string) ([]*schema.GetGoodsResponseData, error) {
	log.Printf("Start to call goods service for GetGoodsDetail, goodsId %s", goodsId)
	defer log.Println("End call goods service for GetGoodsDetail")

	if goodsId == "" {
		return nil, errors.New("[BFF-Adapter-GoodsServiceAdapter-GetGoodsDetail] goodsId must not be empty")
	}

	var result *schema.GetGoodsResponse
	url := fmt.Sprintf("http://localhost:%d/api/goods-service/goods/%s", a.port, goodsId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDetail-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDetail-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDetail-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDetail-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *goodsServiceAdapter) CheckWarehouse(ctx context.Context, request *schema.CheckWarehouseRequest) (*schema.CheckWarehouseResponseData, error) {
	log.Println("Start to call goods service for CheckWarehouse")
	defer log.Println("End call goods service for CheckWarehouse")

	var result *schema.CheckWarehouseResponse
	url := fmt.Sprintf("http://localhost:%d/api/goods-service/check-wh", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CheckWarehouse-Marshal error %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CheckWarehouse-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CheckWarehouse-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CheckWarehouse-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CheckWarehouse-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *goodsServiceAdapter) CreateWHTransfer(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error {
	log.Println("Start to call goods service for CreateWHTransfer")
	defer log.Println("End call goods service for CreateWHTransfer")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/goods-service/wh-transfer", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CreateWHTransfer-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CreateWHTransfer-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CreateWHTransfer-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CreateWHTransfer-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CreateWHTransfer-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
