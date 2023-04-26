package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/order_service/config"
	"store-bpel/order_service/internal/controller"
	"store-bpel/order_service/schema"
	"strconv"
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

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), r); err != nil {
		log.Fatal(err)
	}
	log.Printf("Order Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	// Customer
	r.HandleFunc("/api/order-service/customer/make-order", handleMakeOnlineOrder)
	r.HandleFunc("/api/order-service/customer/{customerId}", handleGetListOrderCustomer)
	r.HandleFunc("/api/order-service/customer/order-detail/{orderId}", handleCustomerGetOrderDetail)
	r.HandleFunc("/api/order-service/online-order-status/{orderId}", handleGetOnlineOrdersState)

	// Admin
	r.HandleFunc("/api/order-service/admin/order-detail/{orderId}", handleAdminGetOrderDetail)
	r.HandleFunc("/api/order-service/admin/order", handleAdminMakeOfflineOrder)
	r.HandleFunc("/api/order-service/admin/offline-order", handleAdminGetOfflineOrder)
	r.HandleFunc("/api/order-service/admin/online-order", handleAdminGetOnlineOrder)

	// Shared
	r.HandleFunc("/api/order-service/ship-fee", handleGetShipFee)
	r.HandleFunc("/api/order-service/online-order-status", handleUpdateOnlineOrdersState)
	r.HandleFunc("/api/order-service/best-goods", handleGetBestSellingGoods)
}

func handleMakeOnlineOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodPost {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.MakeOnlineOrderRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.CreateOnlineOrder(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAdminMakeOfflineOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodPost {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.MakeOfflineOrderRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.CreateOfflineOrder(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetListOrderCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		customerId := mux.Vars(r)["customerId"]
		order, err := ctrl.GetListOrderCustomer(ctx, customerId)
		if err != nil {
			err = enc.Encode(&schema.GetListOrderCustomerResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetListOrderCustomerResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       order,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleCustomerGetOrderDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		orderId := mux.Vars(r)["orderId"]
		order, err := ctrl.GetOrderDetail(ctx, orderId)
		if err != nil {
			err = enc.Encode(&schema.GetOrderDetailCustomerResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetOrderDetailCustomerResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       order,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAdminGetOrderDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		orderId := mux.Vars(r)["orderId"]
		orderIdInt, err := strconv.Atoi(orderId)
		if err != nil {
			err = enc.Encode(&schema.GetOrderDetailAdminResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		}
		order, err := ctrl.GetOrderDetailAdmin(ctx, orderIdInt)
		if err != nil {
			err = enc.Encode(&schema.GetOrderDetailAdminResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetOrderDetailAdminResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       order,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAdminGetOfflineOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		orders, err := ctrl.GetOfflineOrders(ctx)
		if err != nil {
			err = enc.Encode(&schema.GetOfflineOrdersResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetOfflineOrdersResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       orders,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAdminGetOnlineOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		orders, err := ctrl.GetOnlineOrders(ctx)
		if err != nil {
			err = enc.Encode(&schema.GetOnlineOrdersResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetOnlineOrdersResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       orders,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetShipFee(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.GetShipFeeResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.GetShipFeeRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.GetShipFeeResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		fee, err := ctrl.GetShipFee(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.GetShipFeeResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetShipFeeResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       fee,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetOnlineOrdersState(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		orderId := mux.Vars(r)["orderId"]
		status, err := ctrl.GetOnlineOrdersStatus(ctx, orderId)
		if err != nil {
			err = enc.Encode(&schema.GetOnlineOrdersStatusResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetOnlineOrdersStatusResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       status,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateOnlineOrdersState(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodPut {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.UpdateOnlineOrdersStatusRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.UpdateOrderState(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetBestSellingGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		goods, err := ctrl.GetBestSellingGoods(ctx)
		if err != nil {
			err = enc.Encode(&schema.GetBestSellingGoodsResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetBestSellingGoodsResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       goods,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
