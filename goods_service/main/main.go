package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"io/ioutil"
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
	r.HandleFunc("/api/goods-service/default-goods", handleGoodsDefault)
	r.HandleFunc("/api/goods-service/goods:search", handleSearchGoods)
	r.HandleFunc("/api/goods-service/product/{productId}", handleProductDetail)
	r.HandleFunc("/api/goods-service/check-wh", handleCheckWH)
	r.HandleFunc("/api/goods-service/image", handleUploadImage)
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

func handleGoodsDefault(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == "GET" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.GetGoodsDefaultResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.GetGoodsDefaultRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.GetGoodsDefaultResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		resp, err := ctrl.GetGoodsDefault(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.GetGoodsDefaultResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetGoodsDefaultResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleSearchGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == "GET" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.SearchGoodsResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.SearchGoodsRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.SearchGoodsResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		resp, err := ctrl.SearchGoods(ctx, request)
		if err != nil {
			err = enc.Encode(&schema.SearchGoodsResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.SearchGoodsResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleProductDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	productId := mux.Vars(r)["productId"]
	if r.Method == "GET" {
		resp, err := ctrl.GetProductDetail(ctx, productId)
		if err != nil {
			err = enc.Encode(&schema.GetDetailProductsResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetDetailProductsResponse{
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
		var request []*schema.AddGoodsRequest
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
			err = enc.Encode(&schema.GetGoodsDetailResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetGoodsDetailResponse{
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
		var request []*schema.UpdateGoodsRequest
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

// use for debugging internal service
//func handleUploadImage(w http.ResponseWriter, r *http.Request) {
//	ctx := context.Background()
//	w.Header().Set("Content-Type", "application/json")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	enc := json.NewEncoder(w)
//	if r.Method == "POST" {
//		var (
//			goodsId    = r.FormValue("goodsId")
//			goodsColor = r.FormValue("goodsColor")
//			isDefault  = r.FormValue("isDefault")
//		)
//		r.Body = http.MaxBytesReader(w, r.Body, common.MAX_UPLOAD_SIZE) // max size 1MB
//		if err := r.ParseMultipartForm(common.MAX_UPLOAD_SIZE); err != nil {
//			http.Error(w, "Max upload size is 1MB", http.StatusBadRequest)
//			return
//		}
//
//		file, fileHeader, err := r.FormFile("images")
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		defer file.Close()
//
//		// Create the uploads folder if it doesn't
//		// already exist
//		err = os.MkdirAll(fmt.Sprintf("../uploads/%s", goodsId), os.ModePerm)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		// Create a new file in the uploads directory
//		relativePath := fmt.Sprintf("../uploads/%s/%s", goodsId, fileHeader.Filename)
//		dst, err := os.Create(relativePath)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		defer dst.Close()
//
//		// Copy the uploaded file to the filesystem
//		// at the specified destination
//		_, err = io.Copy(dst, file)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		absPath, err := filepath.Abs(relativePath)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		err = ctrl.UploadGoodsImage(ctx, goodsId, goodsColor, absPath, isDefault == "true")
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		_ = enc.Encode(&schema.UpdateResponse{
//			StatusCode: 200,
//			Message:    "OK",
//		})
//	} else {
//		http.Error(w, "Method not supported", http.StatusNotFound)
//	}
//}

func handleUploadImage(w http.ResponseWriter, r *http.Request) {
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
		var request *schema.UploadImageRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.UploadGoodsImage(ctx, request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_ = enc.Encode(&schema.UpdateResponse{
			StatusCode: 200,
			Message:    "OK",
		})
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
