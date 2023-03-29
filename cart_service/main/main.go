package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/cart_service/config"
	"store-bpel/cart_service/controller"
	"store-bpel/cart_service/schema"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
)

var ctrl controller.ICartServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Cart Service server started at port %d", cfg.HttpPort)

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
	log.Printf("Cart Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	r.HandleFunc("/api/cart-service/cart", handleCart)
	r.HandleFunc("/api/cart-service/cart/{customerId}", handleCustomer)
}

func handleCart(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}

		var request *schema.AddCartRequest
		err = json.Unmarshal(reqBody, &request)

		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}

		err = ctrl.AddCart(ctx, request.CustomerId)
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

func handleCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	vars := mux.Vars(r)

	if r.Method == "GET" {
		cart, err := ctrl.GetCart(ctx, vars["customerId"])
		if err != nil {
			err = enc.Encode(&schema.GetCartResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetCartResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       cart,
			})
		}

	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
