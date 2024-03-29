package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"store-bpel/bff/admin_bff/internal/controller"
	account_controller "store-bpel/bff/admin_bff/internal/controller/account_service"
	branch_controller "store-bpel/bff/admin_bff/internal/controller/branch_service"
	event_controller "store-bpel/bff/admin_bff/internal/controller/event_service"
	goods_controller "store-bpel/bff/admin_bff/internal/controller/goods_service"
	order_controller "store-bpel/bff/admin_bff/internal/controller/order_service"
	staff_controller "store-bpel/bff/admin_bff/internal/controller/staff_service"
	statistic_controller "store-bpel/bff/admin_bff/internal/controller/statistic_service"
	warehouse_controller "store-bpel/bff/admin_bff/internal/controller/warehouse_service"

	"log"
	"net/http"
	"store-bpel/bff/admin_bff/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Admin BFF server started at port %d", cfg.HttpPort)

	r := newSOAPMux(cfg)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), r); err != nil {
		log.Fatal(err)
	}
	log.Printf("Admin BFF initialized successfully at port %d", cfg.HttpPort)
}

func newSOAPMux(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	// middleware
	r.Use(controller.AuthMiddleware)

	branch_controller.RegisterEndpointHandler(r, cfg)
	account_controller.RegisterEndpointHandler(r, cfg)
	staff_controller.RegisterEndpointHandler(r, cfg)
	goods_controller.RegisterEndpointHandler(r, cfg)
	event_controller.RegisterEndpointHandler(r, cfg)
	warehouse_controller.RegisterEndpointHandler(r, cfg)
	order_controller.RegisterEndpointHandler(r, cfg)
	statistic_controller.RegisterEndpointHandler(r, cfg)
	return r
}
