package main

import (
	"github.com/spf13/cast"
	"log"
	"net/http"
	"store-bpel/bff/customer_bff/config"
	customer_controller "store-bpel/bff/customer_bff/controller/customer_service"
	order_controller "store-bpel/bff/customer_bff/controller/order_service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Customer BFF server started at port %d", cfg.HttpPort)

	mux := newSOAPMux(cfg)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), mux); err != nil {
		log.Fatal(err)
	}
	log.Printf("Customer BFF initialized successfully at port %d", cfg.HttpPort)
}

func newSOAPMux(cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	customer_controller.RegisterEndpointHandler(mux, cfg)
	order_controller.RegisterEndpointHandler(mux, cfg)
	return mux
}
