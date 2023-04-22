package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/statistic_service/config"
	"store-bpel/statistic_service/controller"
	"store-bpel/statistic_service/schema"
)

var ctrl controller.IStatisticServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Statistic Service server started at port %d", cfg.HttpPort)

	db, err := DbConnect(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ctrl = controller.NewController(cfg, db)

	ctx := context.Background()
	go Consume(ctx, ctrl)

	r := mux.NewRouter()
	registerEndpoint(r)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), r); err != nil {
		log.Fatal(err)
	}
	log.Printf("Statistic Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	r.HandleFunc("/api/statistic-service/overall-stat", handleGetOverallStat)
	r.HandleFunc("/api/statistic-service/revenue", handleGetRevenue)
	r.HandleFunc("/api/statistic-service/revenue/{goodsId}", handleGetRevenueOneGoods)
	r.HandleFunc("/api/statistic-service/profit", handleGetProfit)
	r.HandleFunc("/api/statistic-service/profit/{goodsId}", handleGetProfitOneGoods)
	r.HandleFunc("/api/statistic-service/order", handleAddOrderData)
}

func handleAddOrderData(w http.ResponseWriter, r *http.Request) {
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
		var request *schema.AddOrderDataRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.AddOrderData(ctx, request)
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

func handleGetOverallStat(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.GetOverallStatisticResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.CommonGetStatisticRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.GetOverallStatisticResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		stat, err := ctrl.GetOverallStat(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.GetOverallStatisticResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetOverallStatisticResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       stat,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.GetRevenueResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.FilterGetStatisticRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.GetRevenueResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		revenue, err := ctrl.GetRevenue(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.GetRevenueResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetRevenueResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       revenue,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetRevenueOneGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.GetRevenueResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.CommonGetStatisticRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.GetRevenueResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		goodsId := mux.Vars(r)["goodsId"]
		revenue, err := ctrl.GetRevenueOneGoods(ctx, request, goodsId)
		if err != nil {
			err = enc.Encode(&schema.GetRevenueResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetRevenueResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       revenue,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetProfit(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.GetProfitResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.FilterGetStatisticRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.GetProfitResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		profit, err := ctrl.GetProfit(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.GetProfitResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetProfitResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       profit,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetProfitOneGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.GetProfitResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.CommonGetStatisticRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.GetProfitResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		goodsId := mux.Vars(r)["goodsId"]
		profit, err := ctrl.GetProfitOneGoods(ctx, request, goodsId)
		if err != nil {
			err = enc.Encode(&schema.GetProfitResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetProfitResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       profit,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
