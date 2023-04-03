package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/warehouse_service/config"
	"store-bpel/warehouse_service/controller"
	"store-bpel/warehouse_service/schema"
)

var ctrl controller.IWarehouseServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Warehouse Service server started at port %d", cfg.HttpPort)

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
	log.Printf("Warehouse Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	r.HandleFunc("/api/warehouse-service/manager/{warehouseId}", handleGetWarehouseManager)
	r.HandleFunc("/api/warehouse-service/manager", handleUpdateManager)
	r.HandleFunc("/api/warehouse-service/staff/{warehouseId}", handleWarehouseStaff)
	r.HandleFunc("/api/warehouse-service/staff", handleUpdateStaff)
	r.HandleFunc("/api/warehouse-service/warehouse/{warehouseId}", handleGetWarehouse)
	r.HandleFunc("/api/warehouse-service/warehouse", handleUpdateWarehouse)
}

func handleGetWarehouseManager(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		warehouseId := mux.Vars(r)["warehouseId"]
		resp, err := ctrl.GetWarehouseManager(ctx, warehouseId)
		if err != nil {
			err = enc.Encode(&schema.GetWarehouseManagerResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetWarehouseManagerResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateManager(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&schema.UpdateResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
		return
	}
	if r.Method == http.MethodPut {
		var request *schema.UpdateManagerRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.UpdateWarehouseManager(ctx, request)
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

func handleWarehouseStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		warehouseId := mux.Vars(r)["warehouseId"]
		resp, err := ctrl.GetWarehouseStaff(ctx, warehouseId)
		if err != nil {
			err = enc.Encode(&schema.GetWarehouseStaffResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetWarehouseStaffResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&schema.UpdateResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request *schema.AddWarehouseStaffRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.AddWarehouseStaff(ctx, request)
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
	} else if r.Method == http.MethodPut {
		var request *schema.UpdateWarehouseStaffRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.UpdateStaff(ctx, request)
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
	} else if r.Method == http.MethodDelete {
		var request *schema.DeleteWarehouseStaffRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.DeleteStaff(ctx, request)
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

func handleGetWarehouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
		warehouseId := mux.Vars(r)["warehouseId"]
		resp, err := ctrl.GetWarehouse(ctx, warehouseId)
		if err != nil {
			err = enc.Encode(&schema.GetWarehouseResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetWarehouseResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&schema.UpdateResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request *schema.AddWarehouseRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.AddWarehouse(ctx, request)
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
	} else if r.Method == http.MethodPut {
		var request *schema.UpdateWarehouseRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.UpdateWarehouse(ctx, request)
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
	} else if r.Method == http.MethodDelete {
		var request *schema.DeleteWarehouseRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.DeleteWarehouse(ctx, request)
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
