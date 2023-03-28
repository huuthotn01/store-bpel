package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"store-bpel/event_service/config"
	"store-bpel/event_service/controller"
	"store-bpel/event_service/schema"

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
	r.HandleFunc("/api/event-service/event", handleGetEvent)
}

func handleGetEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------Coming--------")
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
		http.Error(w, "Method not supported", http.StatusNotFound)
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
