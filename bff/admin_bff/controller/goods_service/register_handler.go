package goods_service

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/goods_service"
	"store-bpel/goods_service/common"
	"time"

	"github.com/gorilla/mux"
)

var goodsController IGoodsBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	goodsController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/goods-service/goods", handleGetGoods)
	mux.HandleFunc("/api/bff/goods-service/goods-detail", handleGetGoodsDetail)
	mux.HandleFunc("/api/bff/goods-service/goods/add", handleAddGoods)
	mux.HandleFunc("/api/bff/goods-service/goods/import", handleImport)
	mux.HandleFunc("/api/bff/goods-service/goods/export", handleExport)
	mux.HandleFunc("/api/bff/goods-service/goods/return-manufact", handleReturnManufacturer)
	mux.HandleFunc("/api/bff/goods-service/goods/cust-return", handleCustReturn)
	mux.HandleFunc("/api/bff/goods-service/goods/get-warehouse", handleGetWarehouse)
	mux.HandleFunc("/api/bff/goods-service/goods/update", handleUpdateGoods)
	mux.HandleFunc("/api/bff/goods-service/goods/image", handleUploadImage)
	mux.HandleFunc("/api/bff/goods-service/goods/image:delete", handleDeleteImage)
}

func handleGetGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	if r.Method == http.MethodPost {
		resp, err := goodsController.GetGoods(ctx)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleGetGoods-GetGoods err %v", err),
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

func handleGetGoodsDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleGetGoodsDetail-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.GetGoodsDetailRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleGetGoodsDetail-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := goodsController.GetGoodsDetail(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleGetGoodsDetail-GetCustomer err %v", err),
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

func handleUploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if r.Method == "POST" {
		var (
			goodsId    = r.FormValue("goodsId")
			goodsColor = r.FormValue("goodsColor")
			isDefault  = r.FormValue("isDefault")
		)
		r.Body = http.MaxBytesReader(w, r.Body, common.MAX_UPLOAD_SIZE) // max size 1MB
		if err := r.ParseMultipartForm(common.MAX_UPLOAD_SIZE); err != nil {
			http.Error(w, "Max upload size is 1MB", http.StatusBadRequest)
			return
		}

		file, fileHeader, err := r.FormFile("images")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create the uploads folder if it doesn't
		// already exist
		err = os.MkdirAll(fmt.Sprintf("../uploads/%s", goodsId), os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a new file in the uploads directory
		relativePath := fmt.Sprintf("uploads/%s/%d%s", goodsId, time.Now().Unix(), filepath.Ext(fileHeader.Filename))
		dst, err := os.Create("../" + relativePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer dst.Close()

		// Copy the uploaded file to the filesystem
		// at the specified destination
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		imgPath := "/store-bpel/bff/admin_bff/" + relativePath
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = goodsController.UploadImage(ctx, goodsId, goodsColor, imgPath, isDefault == "true")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_ = enc.Encode(&goods_service.UpdateResponse{
			StatusCode: 200,
			Message:    "OK",
		})
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleDeleteImage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleDeleteImage-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.DeleteImageRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleDeleteImage-xml.Unmarshal err %v", err),
			})
		}
		err = goodsController.DeleteImage(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleDeleteImage-AddGoods err %v", err),
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
		fmt.Println("-------", request)
		err = goodsController.AddGoods(ctx, request.Element)
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

func handleGetWarehouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleGetWarehouse-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.GetWarehouseByGoodsRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleGetWarehouse-xml.Unmarshal err %v", err),
			})
			return
		}
		resp, err := goodsController.GetWarehouseByGoods(ctx, request)
		if err != nil {
			err = enc.Encode(&goods_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleGetWarehouse-GetCustomer err %v", err),
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

func handleUpdateGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&goods_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Goods-handleUpdateGoods-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(goods_service.UpdateGoodsRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleUpdateGoods-xml.Unmarshal err %v", err),
			})
		}
		err = goodsController.UpdateGoods(ctx, request.Element)
		if err != nil {
			err = enc.Encode(&goods_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Goods-handleUpdateGoods-AddGoods err %v", err),
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
