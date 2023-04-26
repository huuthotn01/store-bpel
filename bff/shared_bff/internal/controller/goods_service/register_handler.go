package goods_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/schema/goods_service"
)

var goodsController IGoodsBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	goodsController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/goods-service/goods-default", handleGetGoodsDefault)
	mux.HandleFunc("/api/bff/goods-service/products-detail", handleGetProductsDetail)
	mux.HandleFunc("/api/bff/goods-service/check-wh", handleCheckWarehouse)
	mux.HandleFunc("/api/bff/goods-service/wh-transfer", handleCreateTransfer)
	mux.HandleFunc("/api/bff/goods-service/goods-search", handleSearchGoods)
}

func handleGetGoodsDefault(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleGetGoodsDefault-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.GetGoodsDefaultRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleGetGoodsDefault-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := goodsController.GetGoodsDefault(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleGetGoodsDefault-GetGoodsDefault err %v", err),
			})
		} else {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetProductsDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleGetProductsDetail-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.GetProductsDetailRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleGetProductsDetail-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := goodsController.GetProductsDetail(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleGetProductsDetail-GetProductsDetail err %v", err),
			})
		} else {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleCheckWarehouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleCheckWarehouse-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.CheckWarehouseRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleCheckWarehouse-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := goodsController.CheckWarehouse(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleCheckWarehouse-CheckWarehouse err %v", err),
			})
		} else {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleCreateTransfer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Customer-handleCreateTransfer-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.CreateGoodsTransactionRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleCreateTransfer-xml.Unmarshal err %v", err),
			})
			return
		}
		err = goodsController.CreateTransfer(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleCreateTransfer-CreateTransfer err %v", err),
			})
		} else {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleSearchGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleSearchGoods-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.SearchGoodsRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleSearchGoods-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := goodsController.SearchGoods(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleSearchGoods-CheckWarehouse err %v", err),
			})
		} else {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
