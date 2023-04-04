package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/controller"
	customer_controller "store-bpel/bff/customer_bff/controller/customer_service"
	order_controller "store-bpel/bff/customer_bff/controller/order_service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Customer BFF server started at port %d", cfg.HttpPort)

	r := newSOAPMux(cfg)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), r); err != nil {
		log.Fatal(err)
	}
	log.Printf("Customer BFF initialized successfully at port %d", cfg.HttpPort)
}

func newSOAPMux(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	// middleware
	r.Use(controller.AuthMiddleware)

	customer_controller.RegisterEndpointHandler(r, cfg)
	order_controller.RegisterEndpointHandler(r, cfg)
	return r
}
