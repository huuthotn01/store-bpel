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
	"store-bpel/statistic_service/schema"
)

type IStatisticServiceAdapter interface {
	GetOverallStat(ctx context.Context, request *schema.CommonGetStatisticRequest) ([]*schema.GetOverallStatisticResponseData, error)
	GetRevenue(ctx context.Context, request *schema.FilterGetStatisticRequest) ([]*schema.GetRevenueResponseData, error)
	GetRevenueOneGoods(ctx context.Context, request *schema.CommonGetStatisticRequest, goodsId string) ([]*schema.GetRevenueResponseData, error)
	GetProfit(ctx context.Context, request *schema.FilterGetStatisticRequest) ([]*schema.GetProfitResponseData, error)
	GetProfitOneGoods(ctx context.Context, request *schema.CommonGetStatisticRequest, goodsId string) ([]*schema.GetProfitResponseData, error)
	AddOrderData(ctx context.Context, request *schema.AddOrderDataRequest) error
}

type statisticServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewStatisticAdapter(cfg *config.Config) IStatisticServiceAdapter {
	return &statisticServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.StatisticServicePort,
	}
}

func (a *statisticServiceAdapter) AddOrderData(ctx context.Context, request *schema.AddOrderDataRequest) error {
	log.Println("Start to call statistic service for AddOrderData")
	defer log.Println("End call statistic service for AddOrderData")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/statistic-service/order", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-AddOrderData-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-AddOrderData-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-AddOrderData-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-AddOrderData-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-AddOrderData-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *statisticServiceAdapter) GetOverallStat(ctx context.Context, request *schema.CommonGetStatisticRequest) ([]*schema.GetOverallStatisticResponseData, error) {
	log.Println("Start to call statistic service for GetOverallStat")
	defer log.Println("End call statistic service for GetOverallStat")

	var result *schema.GetOverallStatisticResponse
	url := fmt.Sprintf("http://localhost:%d/api/statistic-service/overall-stat", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetOverallStat-Marshal error %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetOverallStat-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetOverallStat-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetOverallStat-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetOverallStat-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *statisticServiceAdapter) GetRevenue(ctx context.Context, request *schema.FilterGetStatisticRequest) ([]*schema.GetRevenueResponseData, error) {
	log.Println("Start to call statistic service for GetRevenue")
	defer log.Println("End call statistic service for GetRevenue")

	var result *schema.GetRevenueResponse
	url := fmt.Sprintf("http://localhost:%d/api/statistic-service/revenue", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenue-Marshal error %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenue-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenue-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenue-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenue-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *statisticServiceAdapter) GetRevenueOneGoods(ctx context.Context, request *schema.CommonGetStatisticRequest, goodsId string) ([]*schema.GetRevenueResponseData, error) {
	if goodsId == "" {
		return nil, errors.New("[BFF-Adapter-StatisticServiceAdapter-GetRevenueOneGoods] goodsId must not be empty")
	}

	log.Println("Start to call statistic service for GetRevenueOneGoods")
	defer log.Println("End call statistic service for GetRevenueOneGoods")

	var result *schema.GetRevenueResponse
	url := fmt.Sprintf("http://localhost:%d/api/statistic-service/revenue/%s", a.port, goodsId)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenueOneGoods-Marshal error %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenueOneGoods-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenueOneGoods-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenueOneGoods-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetRevenueOneGoods-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *statisticServiceAdapter) GetProfit(ctx context.Context, request *schema.FilterGetStatisticRequest) ([]*schema.GetProfitResponseData, error) {
	log.Println("Start to call statistic service for GetProfit")
	defer log.Println("End call statistic service for GetProfit")

	var result *schema.GetProfitResponse
	url := fmt.Sprintf("http://localhost:%d/api/statistic-service/profit", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfit-Marshal error %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfit-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfit-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfit-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfit-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *statisticServiceAdapter) GetProfitOneGoods(ctx context.Context, request *schema.CommonGetStatisticRequest, goodsId string) ([]*schema.GetProfitResponseData, error) {
	if goodsId == "" {
		return nil, errors.New("[BFF-Adapter-StatisticServiceAdapter-GetRevenueOneGoods] goodsId must not be empty")
	}

	log.Println("Start to call statistic service for GetProfitOneGoods")
	defer log.Println("End call statistic service for GetProfitOneGoods")

	var result *schema.GetProfitResponse
	url := fmt.Sprintf("http://localhost:%d/api/statistic-service/profit/%s", a.port, goodsId)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfitOneGoods-Marshal error %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfitOneGoods-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfitOneGoods-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfitOneGoods-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StatisticServiceAdapter-GetProfitOneGoods-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}
