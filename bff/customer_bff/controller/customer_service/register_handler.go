package customer_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/schema/customer_service"
)

var customerController ICustomerBffController

func RegisterEndpointHandler(mux *http.ServeMux, cfg *config.Config) {
	// init controller
	customerController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/customer-service/customer/get", handleGetCustomer)
	mux.HandleFunc("/api/bff/customer-service/customer/add", handleAddCustomer)
	mux.HandleFunc("/api/bff/customer-service/customer/update", handleUpdateCustomer)
}

func handleGetCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&customer_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Customer-handleGetCustomer-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(customer_service.GetCustomerInfoRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&customer_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleGetCustomer-xml.Unmarshal err %v", err),
			})
			return
		}
		customer, err := customerController.GetCustomer(ctx, request)
		if err != nil {
			err = enc.Encode(&customer_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleGetCustomer-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&customer_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       customer,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAddCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&customer_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Customer-handleAddCustomer-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(customer_service.AddCustomerRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleAddCustomer-xml.Unmarshal err %v", err),
			})
			return
		}
		err = customerController.AddCustomer(ctx, request)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleAddCustomer-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&customer_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Customer-handleUpdateCustomer-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(customer_service.UpdateCustomerInfoRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleUpdateCustomer-xml.Unmarshal err %v", err),
			})
			return
		}
		err = customerController.UpdateCustomer(ctx, request)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleUpdateCustomer-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
