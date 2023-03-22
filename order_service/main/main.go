package main

import (
	"github.com/spf13/cast"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"store-bpel/order_service/config"
	"store-bpel/order_service/controller"
)

var ctrl controller.IOrderServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Order Service server started at port %d", cfg.HttpPort)

	db, err := DbConnect(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ctrl = controller.NewController(cfg, db)

	r := mux.NewRouter()
	registerEndpoint(r)

	if err = http.ListenAndServe(":" + cast.ToString(cfg.HttpPort), r); err != nil {
		log.Fatal(err)
	}
	log.Printf("Order Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	// r.HandleFunc({api}, {handleFunc})
}
