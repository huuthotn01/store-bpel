package event_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/schema/event_service"

	"github.com/gorilla/mux"
)

var eventController IEventBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	eventController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/event-service/event-detail", handleEventDetail)
	mux.HandleFunc("/api/bff/event-service/all-event", handleAllEvent)
	mux.HandleFunc("/api/bff/event-service/current-event", handleCurrentEvent)
	mux.HandleFunc("/api/bff/event-service/event-by-goods", handleEventByGoods)
}

func handleEventDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	// read body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&event_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Event-handleEventDetail-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(event_service.GetEventDetailRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleEventDetail-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := eventController.GetEventDetail(ctx, request)
		if err != nil {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleEventDetail-GetEventDetail err %v", err),
			})
		} else {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
func handleCurrentEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	// read body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&event_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Event-handleCurrentEvent-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(event_service.GetEventCurrentRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleCurrentEvent-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := eventController.GetEventCurrent(ctx, request)
		if err != nil {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleCurrentEvent-GetEventCurrent err %v", err),
			})
		} else {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAllEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	if r.Method == http.MethodPost {
		resp, err := eventController.GetEvent(ctx)
		if err != nil {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleAllEvent-GetEvent err %v", err),
			})
		} else {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleEventByGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	// read body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&event_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Event-handleEventByGoods-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(event_service.GetEventByGoodsRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleEventByGoods-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := eventController.GetEventByGoods(ctx, request)
		if err != nil {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleEventByGoods-GetEventDetail err %v", err),
			})
		} else {
			err = enc.Encode(&event_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
