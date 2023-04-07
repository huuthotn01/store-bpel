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
	"store-bpel/bff/customer_bff/config"
	"store-bpel/order_service/schema"
)

type IOrderServiceAdapter interface {
	CreateOnlineOrders(ctx context.Context, request *schema.MakeOnlineOrderRequest) error
	GetOnlineOrdersStatus(ctx context.Context, orderId int) ([]*schema.GetOnlineOrdersStatusResponseData, error)
	GetListOrderCustomer(ctx context.Context, customerId string) ([]*schema.GetListOrderCustomerResponseData, error)
	GetOrderCustomerDetail(ctx context.Context, orderId string) (*schema.GetOrderDetailCustomerResponseData, error)
	GetShippingFee(ctx context.Context, request *schema.GetShipFeeRequest) (*schema.GetShipFeeResponseData, error)
	UpdateOnlineOrdersStatus(ctx context.Context, request *schema.UpdateOnlineOrdersStatusRequest) error
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

func (a *orderServiceAdapter) GetOnlineOrdersStatus(ctx context.Context, orderId int) ([]*schema.GetOnlineOrdersStatusResponseData, error) {
	log.Printf("Start to call order service for GetOnlineOrdersStatus, orderId %s", orderId)
	defer log.Println("End call order service for GetOnlineOrdersStatus")

	var result *schema.GetOnlineOrdersStatusResponse

	url := fmt.Sprintf("http://localhost:%d/api/order-service/online-order-status/%d", a.port, orderId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOnlineOrdersStatus-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOnlineOrdersStatus-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOnlineOrdersStatus-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOnlineOrdersStatus-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *orderServiceAdapter) GetShippingFee(ctx context.Context, request *schema.GetShipFeeRequest) (*schema.GetShipFeeResponseData, error) {
	log.Printf("Start to call order service for GetShippingFee")
	defer log.Println("End call order service for GetShippingFee")

	var result *schema.GetShipFeeResponse

	url := fmt.Sprintf("http://localhost:%d/api/order-service/ship-fee", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetShippingFee-Marshal error %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetShippingFee-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetShippingFee-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetShippingFee-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetShippingFee-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *orderServiceAdapter) GetListOrderCustomer(ctx context.Context, customerId string) ([]*schema.GetListOrderCustomerResponseData, error) {
	if customerId == "" {
		return nil, errors.New("[BFF-Adapter-OrderServiceAdapter-GetListOrderCustomer] customerId must not be empty")
	}

	log.Printf("Start to call order service for GetListOrderCustomer, customerId %s", customerId)
	defer log.Println("End call order service for GetListOrderCustomer")

	var result *schema.GetListOrderCustomerResponse

	url := fmt.Sprintf("http://localhost:%d/api/order-service/customer/%s", a.port, customerId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOnlineOrdersStatus-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetListOrderCustomer-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetListOrderCustomer-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetListOrderCustomer-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *orderServiceAdapter) GetOrderCustomerDetail(ctx context.Context, orderId string) (*schema.GetOrderDetailCustomerResponseData, error) {
	if orderId == "" {
		return nil, errors.New("[BFF-Adapter-OrderServiceAdapter-GetOrderCustomerDetail] orderId must not be empty")
	}

	log.Printf("Start to call order service for GetOrderCustomerDetail, orderId %s", orderId)
	defer log.Println("End call order service for GetOrderCustomerDetail")

	var result *schema.GetOrderDetailCustomerResponse

	url := fmt.Sprintf("http://localhost:%d/api/order-service/customer/order-detail/%s", a.port, orderId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOrderCustomerDetail-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOrderCustomerDetail-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOrderCustomerDetail-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-GetOrderCustomerDetail-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *orderServiceAdapter) UpdateOnlineOrdersStatus(ctx context.Context, request *schema.UpdateOnlineOrdersStatusRequest) error {
	log.Println("Start to call order service for UpdateOnlineOrdersStatus")
	defer log.Println("End call order service for UpdateOnlineOrdersStatus")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/order-service/online-order-status", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-UpdateOnlineOrdersStatus-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-UpdateOnlineOrdersStatus-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-UpdateOnlineOrdersStatus-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-UpdateOnlineOrdersStatus-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-UpdateOnlineOrdersStatus-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *orderServiceAdapter) CreateOnlineOrders(ctx context.Context, request *schema.MakeOnlineOrderRequest) error {
	log.Println("Start to call order service for CreateOnlineOrders")
	defer log.Println("End call order service for CreateOnlineOrders")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/order-service/customer/make-order", a.port)
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOnlineOrders-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOnlineOrders-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOnlineOrders-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOnlineOrders-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-OrderServiceAdapter-CreateOnlineOrders-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
