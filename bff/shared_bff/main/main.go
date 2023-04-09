package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/controller/account_service"
	"store-bpel/bff/shared_bff/controller/event_service"
	"store-bpel/bff/shared_bff/controller/goods_service"
	"store-bpel/bff/shared_bff/controller/order_service"

	"github.com/spf13/cast"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Shared BFF server started at port %d", cfg.HttpPort)

	r := newSOAPMux(cfg)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), r); err != nil {
		log.Fatal(err)
	}
	log.Printf("Shared BFF initialized successfully at port %d", cfg.HttpPort)
}

func newSOAPMux(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	goods_service.RegisterEndpointHandler(r, cfg)
	event_service.RegisterEndpointHandler(r, cfg)
	account_service.RegisterEndpointHandler(r, cfg)
	order_service.RegisterEndpointHandler(r, cfg)
	return r
}
