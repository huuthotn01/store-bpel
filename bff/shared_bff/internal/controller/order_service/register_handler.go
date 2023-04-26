package order_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/schema/order_service"
)

var orderController IOrderBffController

func RegisterEndpointHandler(r *mux.Router, cfg *config.Config) {
	// init controller
	orderController = NewController(cfg)
	// register handler
	r.HandleFunc("/api/bff/order-service/get-ship-fee", handleGetShippingFee)
	r.HandleFunc("/api/bff/order-service/online-order-status/update", handleUpdateOnlineOrdersStatus)
}

func handleGetShippingFee(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&order_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Order-handleGetShippingFee-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(order_service.Address)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetShippingFee-xml.Unmarshal err %v", err),
			})
			return
		}
		shipFee, err := orderController.GetShippingFee(ctx, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetShippingFee-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       shipFee,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateOnlineOrdersStatus(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&order_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Order-handleUpdateOnlineOrdersStatus-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(order_service.UpdateOnlineOrdersStatusRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&order_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleUpdateOnlineOrdersStatus-xml.Unmarshal err %v", err),
			})
			return
		}
		err = orderController.UpdateOnlineOrdersStatus(ctx, request)
		if err != nil {
			err = enc.Encode(&order_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleUpdateOnlineOrdersStatus-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&order_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
