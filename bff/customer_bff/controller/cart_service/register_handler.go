package cart_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/schema/cart_service"

	"github.com/gorilla/mux"
)

var cartController ICartBffController

func RegisterEndpointHandler(r *mux.Router, cfg *config.Config) {
	// init controller
	cartController = NewController(cfg)
	// register handler
	r.HandleFunc("/api/bff/cart-service/get-cart", handleGetCart)
	r.HandleFunc("/api/bff/cart-service/add-goods", handleAddGoods)
}

func handleGetCart(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&cart_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-CartBFF-handleGetCart-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(cart_service.GetCartRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&cart_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-CartBFF-handleGetCart-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := cartController.GetCart(ctx, request.UserId)
		if err != nil {
			err = enc.Encode(&cart_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-CartBFF-handleGetCart-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&cart_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAddGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&cart_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-CartBFF-handleAddGoods-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(cart_service.AddGoodsRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&cart_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-CartBFF-handleAddGoods-xml.Unmarshal err %v", err),
			})
			return
		}
		err := cartController.AddGoods(ctx, request)
		if err != nil {
			err = enc.Encode(&cart_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-CartBFF-handleAddGoods-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&cart_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
