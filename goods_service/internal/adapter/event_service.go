package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/event_service/schema"
	"store-bpel/goods_service/config"
)

type IEventServiceAdapter interface {
	GetEventByGoods(ctx context.Context, goodsId string) ([]*schema.GetEventByGoodsData, error)
}

type eventServiceAdapter struct {
	httpClient *http.Client
	host       string
	port       int
}

func NewEventAdapter(cfg *config.Config) IEventServiceAdapter {
	return &eventServiceAdapter{
		httpClient: &http.Client{},
		host:       cfg.EventServiceHost,
		port:       cfg.EventServicePort,
	}
}

func (a *eventServiceAdapter) GetEventByGoods(ctx context.Context, goodsId string) ([]*schema.GetEventByGoodsData, error) {
	log.Printf("Start to call event service for GetEventDetail, eventId %s", goodsId)
	defer log.Println("End call event service for GetEventDetail")

	if goodsId == "" {
		return nil, errors.New("[BFF-Adapter-EventServiceAdapter-GetEventByGoods] goodsId must not be empty")
	}

	// call http to event service
	var result *schema.GetEventByGoodsResponse
	url := fmt.Sprintf("http://%s:%d/api/event-service/get-by-goods/%s", a.host, a.port, goodsId)

	// create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Goods-Adapter-EventServiceAdapter-GetEventByGoods-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	// send request to event service
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("Goods-Adapter-EventServiceAdapter-GetEventByGoods-httpClient.Do error %v", err)
		return nil, err
	}

	// Read all data response
	// Convert io.Reader (type off http.Response.Body) to []byte
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Goods-Adapter-EventServiceAdapter-GetEventByGoods-ioutil.ReadAll error %v", err)
		return nil, err
	}

	// Convert []byte to JSON
	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("Goods-Adapter-EventServiceAdapter-GetEventByGoods-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}
