package event_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/event_service"
)

var eventController IEventBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	eventController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/event-service/add-event", handleAddEvent)
	mux.HandleFunc("/api/bff/event-service/update-event", handleUpdateEvent)
	mux.HandleFunc("/api/bff/event-service/delete-event", handleDeleteEvent)
}

func handleAddEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	if r.Method == http.MethodPost {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleAddEvent-ioutil.ReadAll err %v", err),
			})
		}

		var req *event_service.AddEventRequest
		err = xml.Unmarshal(payload, &req)

		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleAddEvent-json.Unmarshal err %v", err),
			})
		}
		err = eventController.AddEvent(ctx, req)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleAddEvent-AddEvent err %v", err),
			})
		} else {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	if r.Method == http.MethodPost {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleUpdateEvent-ioutil.ReadAll err %v", err),
			})
		}

		var req *event_service.UpdateEventRequest
		err = xml.Unmarshal(payload, &req)

		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleUpdateEvent-json.Unmarshal err %v", err),
			})
		}
		err = eventController.UpdateEvent(ctx, req)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleUpdateEvent-UpdateEvent err %v", err),
			})
		} else {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	if r.Method == http.MethodPost {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleDeleteEvent-ioutil.ReadAll err %v", err),
			})
		}

		var req *event_service.DeleteEventRequest
		err = xml.Unmarshal(payload, &req)

		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleDeleteEvent-json.Unmarshal err %v", err),
			})
		}
		err = eventController.DeleteEvent(ctx, req)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleDeleteEvent-DeleteEvent err %v", err),
			})
		} else {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
