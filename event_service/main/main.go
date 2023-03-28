package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/event_service/config"
	"store-bpel/event_service/controller"
	"store-bpel/event_service/schema"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
)

var ctrl controller.IEventServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Event Service server started at port %d", cfg.HttpPort)

	db, err := DbConnect(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ctrl = controller.NewController(cfg, db)

	r := mux.NewRouter()
	registerEndpoint(r)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), r); err != nil {
		log.Fatal(err)
	}
	log.Printf("Event Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	r.HandleFunc("/api/event-service/event", handleEvent)
	r.HandleFunc("/api/event-service/event/{eventId}", handleEventDetail)
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == "GET" {
		resp, err := ctrl.GetEvent(ctx)
		if err != nil {
			err = enc.Encode(&schema.GetEventResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetEventResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.AddEventRequest
		err = json.Unmarshal(reqBody, &request)

		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}

		err = ctrl.AddEvent(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}

	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleEventDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	vars := mux.Vars(r)

	if r.Method == "GET" {
		eventId, err := strconv.Atoi(vars["eventId"])
		if err != nil {
			err = enc.Encode(&schema.GetEventDetailResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		}
		resp, err := ctrl.GetEventDetail(ctx, eventId)
		if err != nil {
			err = enc.Encode(&schema.GetEventDetailResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetEventDetailResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else if r.Method == "PUT" {
		eventId, err := strconv.Atoi(vars["eventId"])
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.UpdateEventRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}

		err = ctrl.UpdateEvent(ctx, eventId, request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else if r.Method == "DELETE" {
		eventId, err := strconv.Atoi(vars["eventId"])
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}

		err = ctrl.DeleteEvent(ctx, eventId)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
