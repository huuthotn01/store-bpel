package order_service

import (
	"net/http"
	"store-bpel/bff/customer_bff/config"
)

var orderController IOrderBffController

func RegisterEndpointHandler(mux *http.ServeMux, cfg *config.Config) {
	// init controller
	orderController = NewController(cfg)
	// register handler
	// mux.HandleFunc("/api/bff/customer-service/customer/get", handleGetCustomer)
}
