package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"store-bpel/staff_service/config"
	"store-bpel/staff_service/controller"
	"store-bpel/staff_service/schema"
)

var ctrl controller.IStaffServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Staff Service server started at port %d", cfg.HttpPort)

	db, err := DbConnect(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ctrl = controller.NewController(cfg, db)

	r := mux.NewRouter()
	registerEndpoint(r)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("Staff Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	r.HandleFunc("/api/staff-service/staff", handleStaff)
}

func handleStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if r.Method == "GET" {
		resp, err := ctrl.GetStaff(ctx)
		if err != nil {
			err = enc.Encode(&schema.GetResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetResponse{
				StatusCode: 500,
				Message:    err.Error(),
				Data:       resp,
			})
		}
	} else if r.Method == "POST" {

	} else if r.Method == "PUT" {

	} else if r.Method == "DELETE" {

	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
