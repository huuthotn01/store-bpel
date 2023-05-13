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
	"store-bpel/event_service/schema"
)

type IEventServiceAdapter interface {
	AddEvent(ctx context.Context, data *schema.AddEventRequest) error
	UpdateEvent(ctx context.Context, eventId string, data *schema.UpdateEventRequest) error
	DeleteEvent(ctx context.Context, eventId string) error
	UploadImage(ctx context.Context, data *schema.UploadImageRequest) error
	DeleteImage(ctx context.Context, eventId string) error
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

func (a *eventServiceAdapter) UploadImage(ctx context.Context, data *schema.UploadImageRequest) error {
	log.Printf("Start to call event service for UploadImage")
	defer log.Println("End call event service for UploadImage")

	// call http to event service
	url := fmt.Sprintf("http://%s:%d/api/event-service/image", a.host, a.port)

	payload, err := json.Marshal(data)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UploadImage-json.Marshal error %v", err)
		return err
	}

	// create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UploadImage-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	// send request to event service
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UploadImage-httpClient.Do error %v", err)
		return err
	}

	// Read all data response
	// Convert io.Reader (type off http.Response.Body) to []byte
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UploadImage-ioutil.ReadAll error %v", err)
		return err
	}

	// Convert []byte to JSON
	var result *schema.UpdateResponse
	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UploadImage-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *eventServiceAdapter) DeleteImage(ctx context.Context, eventId string) error {
	if eventId == "" {
		return errors.New("[BFF-Adapter-EventServiceAdapter-DeleteImage] event id must not be empty")
	}
	log.Printf("Start to call event service for DeleteImage")
	defer log.Println("End call event service for DeleteImage")

	// call http to event service
	url := fmt.Sprintf("http://%s:%d/api/event-service/image/%s", a.host, a.port, eventId)

	// create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-DeleteImage-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	// send request to event service
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-DeleteImage-httpClient.Do error %v", err)
		return err
	}

	// Read all data response
	// Convert io.Reader (type off http.Response.Body) to []byte
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-DeleteImage-ioutil.ReadAll error %v", err)
		return err
	}

	// Convert []byte to JSON
	var result *schema.UpdateResponse
	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-DeleteImage-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *eventServiceAdapter) AddEvent(ctx context.Context, data *schema.AddEventRequest) error {
	log.Printf("Start to call event service for AddEvent")
	defer log.Println("End call event service for AddEvent")

	// call http to event service
	url := fmt.Sprintf("http://%s:%d/api/event-service/event", a.host, a.port)

	payload, err := json.Marshal(data)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-AddEvent-json.Marshal error %v", err)
		return err
	}

	// create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-AddEvent-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	// send request to event service
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-AddEvent-httpClient.Do error %v", err)
		return err
	}

	// Read all data response
	// Convert io.Reader (type off http.Response.Body) to []byte
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-AddEvent-ioutil.ReadAll error %v", err)
		return err
	}

	// Convert []byte to JSON
	var result *schema.UpdateResponse
	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-AddEvent-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *eventServiceAdapter) UpdateEvent(ctx context.Context, eventId string, data *schema.UpdateEventRequest) error {
	log.Printf("Start to call event service for UpdateEvent")
	defer log.Println("End call event service for UpdateEvent")

	if eventId == "" {
		return errors.New("[BFF-Adapter-EventServiceAdapter-GetEventByGoods] eventId must not be empty")
	}

	// call http to event service
	url := fmt.Sprintf("http://%s:%d/api/event-service/event/%s", a.host, a.port, eventId)

	payload, err := json.Marshal(data)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UpdateEvent-json.Marshal error %v", err)
		return err
	}

	// create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UpdateEvent-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	// send request to event service
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UpdateEvent-httpClient.Do error %v", err)
		return err
	}

	// Read all data response
	// Convert io.Reader (type off http.Response.Body) to []byte
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UpdateEvent-ioutil.ReadAll error %v", err)
		return err
	}

	// Convert []byte to JSON
	var result *schema.UpdateResponse
	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-UpdateEvent-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *eventServiceAdapter) DeleteEvent(ctx context.Context, eventId string) error {
	log.Printf("Start to call event service for DeleteEvent")
	defer log.Println("End call event service for DeleteEvent")

	if eventId == "" {
		return errors.New("[BFF-Adapter-EventServiceAdapter-DeleteEvent] eventId must not be empty")
	}

	// call http to event service
	url := fmt.Sprintf("http://%s:%d/api/event-service/event/%s", a.host, a.port, eventId)

	// create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-DeleteEvent-NewRequestWithContext error %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	// send request to event service
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-DeleteEvent-httpClient.Do error %v", err)
		return err
	}

	// Read all data response
	// Convert io.Reader (type off http.Response.Body) to []byte
	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-DeleteEvent-ioutil.ReadAll error %v", err)
		return err
	}

	// Convert []byte to JSON
	var result *schema.UpdateResponse
	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-EventServiceAdapter-DeleteEvent-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
