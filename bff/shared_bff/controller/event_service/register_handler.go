package event_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/schema/event_service"
)

var eventController IEventBffController

func RegisterEndpointHandler(mux *http.ServeMux, cfg *config.Config) {
	// init controller
	eventController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/event-service/event-detail", handleEventDetail)
}

func handleEventDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	// read body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&event_service.GetEventDetailResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Event-handleEventDetail-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(event_service.GetEventDetailRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&event_service.GetEventDetailResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleEventDetail-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := eventController.GetEventDetail(ctx, request)
		if err != nil {
			err = enc.Encode(&event_service.GetEventDetailResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleEventDetail-GetEventDetail err %v", err),
			})
		} else {
			err = enc.Encode(&event_service.GetEventDetailResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
