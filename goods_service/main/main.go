package main

import (
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"store-bpel/goods_service/config"
	"store-bpel/goods_service/controller"
	"store-bpel/goods_service/schema"
)

var ctrl controller.IGoodsServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Goods Service server started at port %d", cfg.HttpPort)

	db, err := DbConnect(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ctrl = controller.NewController(cfg, db)
	registerEndpoint()

	if err = http.ListenAndServe(":" + cast.ToString(cfg.HttpPort), nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("Goods Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint() {
	http.HandleFunc("/api/goods-service/goods", handleGoods)
}

func handleGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if r.Method == "GET" {
		resp, err := ctrl.GetGoods(ctx)
		if err != nil {
			err = enc.Encode(&schema.GetGoodsResponse{
				StatusCode: 500,
				Message: err.Error(),
			})
		} else {
			err = enc.Encode(resp)
		}
	} else if r.Method == "POST" {

	} else if r.Method == "PUT"{

	} else if r.Method == "DELETE" {

	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
