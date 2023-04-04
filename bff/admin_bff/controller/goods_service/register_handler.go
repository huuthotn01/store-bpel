package goods_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/goods_service"

	"github.com/gorilla/mux"
)

var goodsController IGoodsBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	goodsController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/goods-service/goods/add", handleAddGoods)
	mux.HandleFunc("/api/bff/goods-service/goods/import", handleImport)
	mux.HandleFunc("/api/bff/goods-service/goods/export", handleExport)
	mux.HandleFunc("/api/bff/goods-service/goods/return-manufact", handleReturnManufacturer)
	mux.HandleFunc("/api/bff/goods-service/goods/cust-return", handleCustReturn)
}

func handleAddGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleAddGoods-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.AddGoodsRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleAddGoods-xml.Unmarshal err %v", err),
			})
		}
		err = goodsController.AddGoods(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleAddGoods-AddGoods err %v", err),
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

func handleImport(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleImport-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.CreateGoodsTransactionRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleImport-xml.Unmarshal err %v", err),
			})
		}
		err = goodsController.Import(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleImport-Import err %v", err),
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

func handleExport(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleExport-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.CreateGoodsTransactionRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleExport-xml.Unmarshal err %v", err),
			})
		}
		err = goodsController.Export(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleExport-Export err %v", err),
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

func handleReturnManufacturer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleReturnManufacturer-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.CreateGoodsTransactionRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleReturnManufacturer-xml.Unmarshal err %v", err),
			})
		}
		err = goodsController.ReturnManufacturer(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleReturnManufacturer-ReturnManufacturer err %v", err),
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

func handleCustReturn(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleCustReturn-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.CreateGoodsTransactionRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleCustReturn-xml.Unmarshal err %v", err),
			})
		}
		err = goodsController.CustomerReturn(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleCustReturn-CustomerReturn err %v", err),
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
