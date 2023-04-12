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
	"store-bpel/bff/admin_bff/config"
	"store-bpel/goods_service/schema"
)

type IGoodsServiceAdapter interface {
	GetGoods(ctx context.Context) ([]*schema.GetGoodsResponseData, error)
	GetGoodsDetail(ctx context.Context, goodsId string) (*schema.GetGoodsResponseData, error)
	AddGoods(ctx context.Context, request []*schema.AddGoodsRequest) error
	Import(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error
	Export(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error
	ReturnManufacturer(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error
	CustomerReturn(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error
	GetWarehouseByGoods(ctx context.Context, goodsId string) ([]*schema.GetGoodsInWarehouseResponseData, error)
	UpdateGoods(ctx context.Context, goodsId string, request []*schema.UpdateGoodsRequest) error
	UploadImage(ctx context.Context, request *schema.UploadImageRequest) error
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

func (a *goodsServiceAdapter) GetGoodsDetail(ctx context.Context, goodsId string) (*schema.GetGoodsResponseData, error) {
	log.Printf("Start to call goods service for GetGoodsDetail, goodsId %s", goodsId)
	defer log.Println("End call goods service for GetGoodsDetail")

	if goodsId == "" {
		return nil, errors.New("[BFF-Adapter-GoodsServiceAdapter-GetGoodsDetail] goodsId must not be empty")
	}

	var result *schema.GetGoodsDetailResponse
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

func (a *goodsServiceAdapter) AddGoods(ctx context.Context, request []*schema.AddGoodsRequest) error {
	log.Printf("Start to call goods service for AddGoods")
	defer log.Println("End call goods service for AddGoods")

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/goods-service/goods", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-AddGoods-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-AddGoods-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-AddGoods-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-AddGoods-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-AddGoods-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *goodsServiceAdapter) Import(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error {
	log.Printf("Start to call goods service for Import")
	defer log.Println("End call goods service for Import")

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/goods-service/import", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Import-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Import-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Import-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Import-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Import-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *goodsServiceAdapter) Export(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error {
	log.Printf("Start to call goods service for Export")
	defer log.Println("End call goods service for Export")

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/goods-service/export", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Export-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Export-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Export-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Export-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-Export-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *goodsServiceAdapter) ReturnManufacturer(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error {
	log.Printf("Start to call goods service for ReturnManufacturer")
	defer log.Println("End call goods service for ReturnManufacturer")

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/goods-service/return-manufact", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-ReturnManufacturer-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-ReturnManufacturer-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-ReturnManufacturer-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-ReturnManufacturer-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-ReturnManufacturer-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *goodsServiceAdapter) CustomerReturn(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error {
	log.Printf("Start to call goods service for CustomerReturn")
	defer log.Println("End call goods service for CustomerReturn")

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/goods-service/cust-return", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CustomerReturn-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CustomerReturn-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CustomerReturn-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CustomerReturn-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-CustomerReturn-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *goodsServiceAdapter) GetWarehouseByGoods(ctx context.Context, goodsId string) ([]*schema.GetGoodsInWarehouseResponseData, error) {
	log.Printf("Start to call goods service for GetWarehouseByGoods, goodsId %s", goodsId)
	defer log.Println("End call goods service for GetWarehouseByGoods")

	if goodsId == "" {
		return nil, errors.New("[BFF-Adapter-GoodsServiceAdapter-GetWarehouseByGoods] goodsId must not be empty")
	}

	var result *schema.GetWarehouseByGoodsResponse
	url := fmt.Sprintf("http://localhost:%d/api/goods-service/goods/warehouse/%s", a.port, goodsId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetWarehouseByGoods-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetWarehouseByGoods-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetWarehouseByGoods-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetWarehouseByGoods-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *goodsServiceAdapter) UpdateGoods(ctx context.Context, goodsId string, request []*schema.UpdateGoodsRequest) error {
	log.Printf("Start to call goods service for UpdateGoods goodsId = %s", goodsId)
	defer log.Println("End call goods service for UpdateGoods")

	if goodsId == "" {
		return errors.New("[BFF-Adapter-GoodsServiceAdapter-UpdateGoods] goodsId must not be empty")
	}

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/goods-service/goods/%s", a.port, goodsId)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UpdateGoods-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UpdateGoods-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UpdateGoods-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UpdateGoods-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UpdateGoods-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *goodsServiceAdapter) UploadImage(ctx context.Context, request *schema.UploadImageRequest) error {
	log.Printf("Start to call goods service for UploadImage")
	defer log.Println("End call goods service for UploadImage")

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/goods-service/image", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UploadImage-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UploadImage-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UploadImage-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UploadImage-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-UploadImage-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
