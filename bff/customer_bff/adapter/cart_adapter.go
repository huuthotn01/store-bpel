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
	"store-bpel/cart_service/schema"
)

type ICartServiceAdapter interface {
	GetCart(ctx context.Context, username string) (*schema.CartData, error)
	AddGoods(ctx context.Context, cartId string, request []*schema.AddGoodsRequest) error
	UpdateGoods(ctx context.Context, cartId string, request []*schema.UpdateGoodsRequest) error
	DeleteGoods(ctx context.Context, cartId string, request []*schema.DeleteGoodsRequest) error
	DeleteAllGoods(ctx context.Context, cartId string) error
}

type cartServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewCartAdapter(cfg *config.Config) ICartServiceAdapter {
	return &cartServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.CartServicePort,
	}
}

func (a *cartServiceAdapter) GetCart(ctx context.Context, username string) (*schema.CartData, error) {
	if username == "" {
		return nil, errors.New("[BFF-Adapter-CartServiceAdapter-GetCart] username must not be empty")
	}

	log.Printf("Start to call cart service for GetCart, username %s", username)
	defer log.Println("End call cart service for GetCart")

	var result *schema.GetCartResponse

	url := fmt.Sprintf("http://localhost:%d/api/cart-service/cart/%s", a.port, username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-GetCart-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-GetCart-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-GetCart-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-GetCart-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *cartServiceAdapter) AddGoods(ctx context.Context, cartId string, request []*schema.AddGoodsRequest) error {
	if cartId == "" {
		return errors.New("[BFF-Adapter-CartServiceAdapter-AddGood] cartId must not be empty")
	}

	log.Printf("Start to call cart service for AddGoods, cartId %s", cartId)
	defer log.Println("End call cart service for AddGoods")

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-AddGood-Marshal error %v", err)
		return err
	}

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/cart-service/goods/%s", a.port, cartId)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-AddGood-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-AddGood-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-AddGood-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-AddGood-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *cartServiceAdapter) UpdateGoods(ctx context.Context, cartId string, request []*schema.UpdateGoodsRequest) error {
	if cartId == "" {
		return errors.New("[BFF-Adapter-CartServiceAdapter-UpdateGoods] cartId must not be empty")
	}

	log.Printf("Start to call cart service for UpdateGoods, cartId %s", cartId)
	defer log.Println("End call cart service for UpdateGoods")

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-UpdateGoods-Marshal error %v", err)
		return err
	}

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/cart-service/goods/%s", a.port, cartId)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-UpdateGoods-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-UpdateGoods-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-UpdateGoods-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-UpdateGoods-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *cartServiceAdapter) DeleteGoods(ctx context.Context, cartId string, request []*schema.DeleteGoodsRequest) error {
	if cartId == "" {
		return errors.New("[BFF-Adapter-CartServiceAdapter-DeleteGoods] cartId must not be empty")
	}

	log.Printf("Start to call cart service for DeleteGoods, cartId %s", cartId)
	defer log.Println("End call cart service for DeleteGoods")

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-DeleteGoods-Marshal error %v", err)
		return err
	}

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/cart-service/goods/%s", a.port, cartId)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-DeleteGoods-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-DeleteGoods-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-DeleteGoods-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-DeleteGoods-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *cartServiceAdapter) DeleteAllGoods(ctx context.Context, cartId string) error {
	if cartId == "" {
		return errors.New("[BFF-Adapter-CartServiceAdapter-DeleteAllGoods] cartId must not be empty")
	}

	log.Printf("Start to call cart service for DeleteAllGoods, cartId %s", cartId)
	defer log.Println("End call cart service for DeleteAllGoods")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/cart-service/all-goods/%s", a.port, cartId)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-DeleteAllGoods-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-DeleteAllGoods-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-DeleteAllGoods-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-CartServiceAdapter-DeleteAllGoods-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
