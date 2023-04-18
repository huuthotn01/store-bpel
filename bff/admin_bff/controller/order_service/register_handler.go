package order_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/order_service"

	"github.com/gorilla/mux"
)

var orderController IOrderBffController

func RegisterEndpointHandler(r *mux.Router, cfg *config.Config) {
	// init controller
	orderController = NewController(cfg)
	// register handler
	r.HandleFunc("/api/bff/order-service/admin/make-order", handleMakeOfflineOrders)
	r.HandleFunc("/api/bff/order-service/admin/get-order-detail", handleGetOrderDetail)
	r.HandleFunc("/api/bff/order-service/admin/get-online-orders", handleGetOnlineOrders)
	r.HandleFunc("/api/bff/order-service/admin/get-offline-orders", handleGetOfflineOrders)
	r.HandleFunc("/api/bff/order-service/admin/customer-order", handleGetListOrderCustomer)
}

func handleMakeOfflineOrders(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&order_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Order-handleMakeOfflineOrders-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(order_service.MakeOfflineOrderRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&order_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleMakeOfflineOrders-xml.Unmarshal err %v", err),
			})
			return
		}
		err = orderController.CreateOfflineOrder(ctx, request)
		if err != nil {
			err = enc.Encode(&order_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleMakeOfflineOrders-GetCustomer err %v", err),
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

func handleGetOrderDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&order_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Order-handleGetOrderDetail-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(order_service.GetOrderDetailRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetOrderDetail-xml.Unmarshal err %v", err),
			})
			return
		}
		order, err := orderController.GetOrderDetail(ctx, request)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetOrderDetail-GetCustomer err %v", err),
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

func handleGetOnlineOrders(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	if r.Method == http.MethodPost {
		orders, err := orderController.GetOnlineOrders(ctx)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetOnlineOrders-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       orders,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetOfflineOrders(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	if r.Method == http.MethodPost {
		orders, err := orderController.GetOfflineOrders(ctx)
		if err != nil {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Order-handleGetOfflineOrders-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&order_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       orders,
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
				Message:    fmt.Sprintf("BFF-Order-handleGetListOrderCustomer-GetListOrderCustomer err %v", err),
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
