package branch_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/branch_service"
)

var branchController IBranchBffController

func RegisterEndpointHandler(mux *http.ServeMux, cfg *config.Config) {
	// init controller
	branchController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/branch-service/get", handleGetBranch)
	mux.HandleFunc("/api/bff/branch-service/add", handleAddBranch)
}

func handleGetBranch(w http.ResponseWriter, r *http.Request) {
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
		branch, err := branchController.GetBranch(ctx, request.BranchId)
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
