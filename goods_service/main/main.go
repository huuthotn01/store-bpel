package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/goods_service/config"
	"store-bpel/goods_service/controller"
	"store-bpel/goods_service/schema"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
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

	r := mux.NewRouter()
	registerEndpoint(r)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), r); err != nil {
		log.Fatal(err)
	}
	log.Printf("Goods Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	r.HandleFunc("/api/goods-service/import", handleImport)
	r.HandleFunc("/api/goods-service/export", handleExport)
	r.HandleFunc("/api/goods-service/wh-transfer", handleTransfer)
	r.HandleFunc("/api/goods-service/return-manufact", handleReturnManufact)
	r.HandleFunc("/api/goods-service/cust-return", handleCustReturn)
	r.HandleFunc("/api/goods-service/goods/{goodsId}", handleDetailGoods)
	r.HandleFunc("/api/goods-service/goods/warehouse/{goodsId}", handleWarehouse)
	r.HandleFunc("/api/goods-service/goods", handleGoods)
	r.HandleFunc("/api/goods-service/check-wh", handleCheckWH)
}

func handleCheckWH(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.CheckWarehouseResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.CheckWarehouseRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.CheckWarehouseResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		resp, err := ctrl.CheckWarehouse(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.CheckWarehouseResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.CheckWarehouseResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == "GET" {
		resp, err := ctrl.GetGoods(ctx)
		if err != nil {
			err = enc.Encode(&schema.GetGoodsResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetGoodsResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.AddGoodsRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.AddGoods(ctx, request)
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

func handleDetailGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	goodsId := mux.Vars(r)["goodsId"]
	if r.Method == "GET" {
		resp, err := ctrl.GetDetailGoods(ctx, goodsId)
		if err != nil {
			err = enc.Encode(&schema.GetDetailGoodsResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetDetailGoodsResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else if r.Method == "DELETE" {
		err := ctrl.DeleteGoods(ctx, goodsId)
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
	} else if r.Method == "PUT" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.UpdateGoodsRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.UpdateGoods(ctx, request, goodsId)
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

func handleImport(w http.ResponseWriter, r *http.Request) {
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
		var request *schema.CreateGoodsTransactionRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.CreateGoodsTransaction(ctx, request, "IMPORT")
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

func handleExport(w http.ResponseWriter, r *http.Request) {
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
		var request *schema.CreateGoodsTransactionRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.CreateGoodsTransaction(ctx, request, "EXPORT")
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

func handleTransfer(w http.ResponseWriter, r *http.Request) {
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
		var request *schema.CreateGoodsTransactionRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.CreateGoodsTransaction(ctx, request, "TRANSFER")
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

func handleReturnManufact(w http.ResponseWriter, r *http.Request) {
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
		var request *schema.CreateGoodsTransactionRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.CreateGoodsTransaction(ctx, request, "RETURN_MANUFACT")
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

func handleCustReturn(w http.ResponseWriter, r *http.Request) {
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
		var request *schema.CreateGoodsTransactionRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.CreateGoodsTransaction(ctx, request, "CUST_RETURN")
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

func handleWarehouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	goodsId := mux.Vars(r)["goodsId"]
	if r.Method == "GET" {
		resp, err := ctrl.GetWarehouseByGoods(ctx, goodsId)
		if err != nil {
			err = enc.Encode(&schema.GetWarehouseByGoodsResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetWarehouseByGoodsResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
