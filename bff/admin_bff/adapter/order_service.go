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
	"store-bpel/order_service/schema"
)

type IOrderServiceAdapter interface {
	CreateOfflineOrders(ctx context.Context, request *schema.MakeOfflineOrderRequest) error
	GetOrderDetail(ctx context.Context, orderId int) (*schema.GetOrderDetailAdminResponseData, error)
	GetOnlineOrders(ctx context.Context) ([]*schema.GetOnlineOrdersResponseData, error)
	GetOfflineOrders(ctx context.Context) ([]*schema.GetOfflineOrdersResponseData, error)
}

type orderServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewOrderAdapter(cfg *config.Config) IOrderServiceAdapter {
	return &orderServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.OrderServicePort,
	}
}

func (a *orderServiceAdapter) GetOrderDetail(ctx context.Context, orderId int) (*schema.GetOrderDetailAdminResponseData, error) {
	log.Printf("Start to call order service for GetOrderDetail, orderId %d", orderId)
	defer log.Println("End call order service for GetOrderDetail")

	var result *schema.GetOrderDetailAdminResponse

	url := fmt.Sprintf("http://localhost:%d/api/order-service/admin/order-detail/%d", a.port, orderId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOrderDetail-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOrderDetail-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOrderDetail-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOrderDetail-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *orderServiceAdapter) CreateOfflineOrders(ctx context.Context, request *schema.MakeOfflineOrderRequest) error {
	log.Println("Start to call order service for CreateOfflineOrders")
	defer log.Println("End call order service for CreateOfflineOrders")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/order-service/admin/order", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOfflineOrders-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOfflineOrders-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOfflineOrders-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOfflineOrders-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOfflineOrders-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *orderServiceAdapter) GetOnlineOrders(ctx context.Context) ([]*schema.GetOnlineOrdersResponseData, error) {
	log.Printf("Start to call order service for GetOnlineOrders")
	defer log.Println("End call order service for GetOnlineOrders")

	var result *schema.GetOnlineOrdersResponse

	url := fmt.Sprintf("http://localhost:%d/api/order-service/admin/online-order", a.port)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOnlineOrders-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOnlineOrders-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOnlineOrders-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOnlineOrders-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *orderServiceAdapter) GetOfflineOrders(ctx context.Context) ([]*schema.GetOfflineOrdersResponseData, error) {
	log.Printf("Start to call order service for GetOfflineOrders")
	defer log.Println("End call order service for GetOfflineOrders")

	var result *schema.GetOfflineOrdersResponse

	url := fmt.Sprintf("http://localhost:%d/api/order-service/admin/offline-order", a.port)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOfflineOrders-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOfflineOrders-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOfflineOrders-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOfflineOrders-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}
