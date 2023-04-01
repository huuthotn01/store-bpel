package main

import (
	"log"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	account_controller "store-bpel/bff/admin_bff/controller/account_service"
	branch_controller "store-bpel/bff/admin_bff/controller/branch_service"
	event_controller "store-bpel/bff/admin_bff/controller/event_service"
	goods_controller "store-bpel/bff/admin_bff/controller/goods_service"
	staff_controller "store-bpel/bff/admin_bff/controller/staff_service"

	"github.com/spf13/cast"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Admin BFF server started at port %d", cfg.HttpPort)

	mux := newSOAPMux(cfg)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), mux); err != nil {
		log.Fatal(err)
	}
	log.Printf("Admin BFF initialized successfully at port %d", cfg.HttpPort)
}

func newSOAPMux(cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	branch_controller.RegisterEndpointHandler(mux, cfg)
	account_controller.RegisterEndpointHandler(mux, cfg)
	staff_controller.RegisterEndpointHandler(mux, cfg)
	goods_controller.RegisterEndpointHandler(mux, cfg)
	event_controller.RegisterEndpointHandler(mux, cfg)
	return mux
}
