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
	"store-bpel/cart_service/schema"
	"store-bpel/customer_service/config"
)

type ICartServiceAdapter interface {
	AddCart(ctx context.Context, customerId string) error
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

func (a *cartServiceAdapter) AddCart(ctx context.Context, customerId string) error {
	log.Println("Start to call cart service for AddCart")
	defer log.Println("End call cart service for AddCart")

	var result *schema.UpdateResponse
	url := fmt.Sprintf("http://localhost:%d/api/cart-service/cart", a.port)

	request := &schema.AddCartRequest{
		CustomerId: customerId,
	}

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("Customer Service-CartServiceAdapter-AddCart-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("Customer Service-CartServiceAdapter-AddCart-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("Customer Service-CartServiceAdapter-AddCart-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Customer Service-CartServiceAdapter-AddCart-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("Customer Service-CartServiceAdapter-AddCart-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
