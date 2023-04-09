package order_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/schema/order_service"
)

var orderController IOrderBffController

func RegisterEndpointHandler(r *mux.Router, cfg *config.Config) {
	// init controller
	orderController = NewController(cfg)
	// register handler
	r.HandleFunc("/api/bff/order-service/customer/make-order", handleMakeOnlineOrders)
	r.HandleFunc("/api/bff/order-service/customer/get-list", handleGetListOrderCustomer)
	r.HandleFunc("/api/bff/order-service/customer/get-detail", handleGetOrderCustomerDetail)
	r.HandleFunc("/api/bff/order-service/online-order-status/get", handleGetOnlineOrdersStatus)
}

func handleMakeOnlineOrders(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&order_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Order-handleMakeOnlineOrders-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(order_service.MakeOnlineOrderRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&order_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleMakeOnlineOrders-xml.Unmarshal err %v", err),
			})
			return
		}
		err = orderController.CreateOnlineOrder(ctx, request)
		if err != nil {
			err = enc.Encode(&order_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleMakeOnlineOrders-GetCustomer err %v", err),
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

func handleGetOnlineOrdersStatus(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&order_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Order-handleGetOnlineOrdersStatus-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(order_service.GetOnlineOrdersStatusRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetOnlineOrdersStatus-xml.Unmarshal err %v", err),
			})
			return
		}
		status, err := orderController.GetOnlineOrdersStatus(ctx, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetOnlineOrdersStatus-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       status,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetListOrderCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&order_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Order-handleGetListOrderCustomer-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(order_service.GetListOrderCustomerRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetListOrderCustomer-xml.Unmarshal err %v", err),
			})
			return
		}
		status, err := orderController.GetListOrderCustomer(ctx, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetListOrderCustomer-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       status,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetOrderCustomerDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&order_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Order-handleGetOrderCustomerDetail-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(order_service.GetOrderDetailCustomerRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetOrderCustomerDetail-xml.Unmarshal err %v", err),
			})
			return
		}
		order, err := orderController.GetOrderCustomerDetail(ctx, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetOrderCustomerDetail-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       order,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
