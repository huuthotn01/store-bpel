package warehouse_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/warehouse_service"

	"github.com/gorilla/mux"
)

var warehouseController IWarehouseBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	warehouseController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/warehouse-service/get-manager", handleGetManager)
	mux.HandleFunc("/api/bff/warehouse-service/update-manager", handleUpdateManager)
	mux.HandleFunc("/api/bff/warehouse-service/get-staff", handleGetStaff)
	mux.HandleFunc("/api/bff/warehouse-service/add-staff", handleAddStaff)
	mux.HandleFunc("/api/bff/warehouse-service/update-staff", handleUpdateStaff)
	mux.HandleFunc("/api/bff/warehouse-service/delete-staff", handleDeleteStaff)
	mux.HandleFunc("/api/bff/warehouse-service/get-warehouse", handleGetWarehouse)
	mux.HandleFunc("/api/bff/warehouse-service/get-all-warehouse", handleGetAllWarehouse)
	mux.HandleFunc("/api/bff/warehouse-service/add-warehouse", handleAddWarehouse)
	mux.HandleFunc("/api/bff/warehouse-service/update-warehouse", handleUpdateWarehouse)
	mux.HandleFunc("/api/bff/warehouse-service/delete-warehouse", handleDeleteWarehouse)
}

func handleGetManager(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&warehouse_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleGetManager-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(warehouse_service.GetWarehouseId)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleGetManager-xml.Unmarshal err %v", err),
			})
		}
		staff, err := warehouseController.GetManager(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleGetManager-GetManager err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       staff,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateManager(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&warehouse_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleUpdateManager-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(warehouse_service.UpdateManagerRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleUpdateManager-xml.Unmarshal err %v", err),
			})
		}
		err = warehouseController.UpdateManager(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleUpdateManager-UpdateManager err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&warehouse_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleGetStaff-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(warehouse_service.GetWarehouseId)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleGetStaff-xml.Unmarshal err %v", err),
			})
		}
		staffs, err := warehouseController.GetStaff(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleGetStaff-GetStaff err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       staffs,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAddStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&warehouse_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleAddStaff-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(warehouse_service.AddWarehouseStaffRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleAddStaff-xml.Unmarshal err %v", err),
			})
		}
		err = warehouseController.AddStaff(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleAddStaff-AddStaff err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&warehouse_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleUpdateStaff-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(warehouse_service.UpdateWarehouseStaffRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleUpdateStaff-xml.Unmarshal err %v", err),
			})
		}
		err = warehouseController.UpdateStaff(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleUpdateStaff-UpdateStaff err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleDeleteStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&warehouse_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleDeleteStaff-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(warehouse_service.DeleteWarehouseStaffRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleDeleteStaff-xml.Unmarshal err %v", err),
			})
		}
		err = warehouseController.DeleteStaff(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleDeleteStaff-DeleteStaff err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.UpdateResponse{
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
		err = enc.Encode(&warehouse_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleGetWarehouse-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(warehouse_service.GetWarehouseId)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleGetWarehouse-xml.Unmarshal err %v", err),
			})
		}
		warehouse, err := warehouseController.GetWarehouse(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleGetWarehouse-GetWarehouse err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       warehouse,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetAllWarehouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	if r.Method == http.MethodPost {
		warehouses, err := warehouseController.GetAllWarehouse(ctx)
		if err != nil {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleGetAllWarehouse-GetWarehouse err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       warehouses,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAddWarehouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&warehouse_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleAddWarehouse-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(warehouse_service.AddWarehouseRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleAddWarehouse-xml.Unmarshal err %v", err),
			})
		}
		err = warehouseController.AddWarehouse(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleAddWarehouse-AddWarehouse err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&warehouse_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleUpdateWarehouse-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(warehouse_service.UpdateWarehouseRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleUpdateWarehouse-xml.Unmarshal err %v", err),
			})
		}
		err = warehouseController.UpdateWarehouse(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleUpdateWarehouse-UpdateWarehouse err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleDeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&warehouse_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Warehouse-handleDeleteWarehouse-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(warehouse_service.DeleteWarehouseRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleDeleteWarehouse-xml.Unmarshal err %v", err),
			})
		}
		err = warehouseController.DeleteWarehouse(ctx, request)
		if err != nil {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Warehouse-handleDeleteWarehouse-DeleteWarehouse err %v", err),
			})
		} else {
			err = enc.Encode(&warehouse_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
