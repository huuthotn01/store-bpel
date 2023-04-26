package branch_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/branch_service"
)

var branchController IBranchBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	branchController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/branch-service/get-all", handleGetBranch)
	mux.HandleFunc("/api/bff/branch-service/get", handleGetBranchDetail)
	mux.HandleFunc("/api/bff/branch-service/add", handleAddBranch)
	mux.HandleFunc("/api/bff/branch-service/update", handleUpdateBranch)
	mux.HandleFunc("/api/bff/branch-service/manager/update", handleUpdateBranchManager)
	mux.HandleFunc("/api/bff/branch-service/delete", handleDeleteBranch)
	mux.HandleFunc("/api/bff/branch-service/staff/get", handleGetBranchStaff)
}

func handleGetBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	if r.Method == http.MethodPost {
		branch, err := branchController.GetBranch(ctx)
		if err != nil {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleGetBranch-GetBranch err %v", err),
			})
		} else {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       branch,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetBranchDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&branch_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Branch-handleGetBranch-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(branch_service.GetBranchDetailRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleGetBranch-xml.Unmarshal err %v", err),
			})
		}
		branch, err := branchController.GetBranchDetail(ctx, request.BranchId)
		if err != nil {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleGetBranch-GetBranch err %v", err),
			})
		} else {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       branch,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAddBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&branch_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Branch-handleAddBranch-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(branch_service.AddBranchRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleAddBranch-xml.Unmarshal err %v", err),
			})
		}
		err = branchController.AddBranch(ctx, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleAddBranch-AddBranch err %v", err),
			})
		} else {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&branch_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Branch-handleUpdateBranch-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(branch_service.UpdateBranchRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleUpdateBranch-xml.Unmarshal err %v", err),
			})
		}
		err = branchController.UpdateBranch(ctx, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleUpdateBranch-UpdateBranch err %v", err),
			})
		} else {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateBranchManager(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&branch_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Branch-handleUpdateBranchManager-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(branch_service.UpdateBranchManagerRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleUpdateBranchManager-xml.Unmarshal err %v", err),
			})
		}
		err = branchController.UpdateBranchManager(ctx, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleUpdateBranchManager-AddBranch err %v", err),
			})
		} else {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleDeleteBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&branch_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Branch-handleDeleteBranch-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(branch_service.DeleteBranchRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleDeleteBranch-xml.Unmarshal err %v", err),
			})
		}
		err = branchController.DeleteBranch(ctx, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleDeleteBranch-AddBranch err %v", err),
			})
		} else {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetBranchStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&branch_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Branch-handleGetBranchStaff-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(branch_service.GetBranchStaffRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleGetBranchStaff-xml.Unmarshal err %v", err),
			})
		}
		staffs, err := branchController.GetBranchStaff(ctx, request)
		if err != nil {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Branch-handleGetBranchStaff-AddBranch err %v", err),
			})
		} else {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       staffs.Staffs,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
