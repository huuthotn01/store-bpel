package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/event_service/schema"
)

type IEventServiceAdapter interface {
	GetEventDetail(ctx context.Context, eventId string) (*schema.GetEventData, error)
}

type eventServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewEventAdapter(cfg *config.Config) IEventServiceAdapter {
	return &eventServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.EventServicePort,
	}
}

func (a *eventServiceAdapter) GetEventDetail(ctx context.Context, eventId string) (*schema.GetEventData, error) {
	log.Printf("Start to call event service for GetEventDetail, eventId %s", eventId)
	defer log.Println("End call event service for GetEventDetail")

	if eventId == "" {
		return nil, errors.New("[BFF-Adapter-EventServiceAdapter-GetEventDetail] eventId must not be empty")
	}

	// call http to event service
	var result *schema.GetEventDetailResponse
	url := fmt.Sprintf("http://localhost:%d/api/event-service/event/%s", a.port, eventId)

	// create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-GetEventDetail-NewRequestWithContext error %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	// send request to event service
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-GetEventDetail-httpClient.Do error %v", err)
		return nil, err
	}

	// Read all data response
	// Convert io.Reader (type off http.Response.Body) to []byte
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-GoodsServiceAdapter-GetGoodsDetail-ioutil.ReadAll error %v", err)
		return nil, err
	}

	// Convert []byte to JSON
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
