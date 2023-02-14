package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/branch_service/config"
	"store-bpel/branch_service/controller"
	"store-bpel/branch_service/schema"
)

var ctrl controller.IBranchServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Branch Service server started at port %d", cfg.HttpPort)

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
	log.Printf("Branch Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	r.HandleFunc("/api/branch-service/{branchId}", handleBranchDetail)
	r.HandleFunc("/api/branch-service", handleBranch)
	r.HandleFunc("/api/branch-service/manager/{branchId}", handleBranchManager)
	r.HandleFunc("/api/branch-service/staff/{branchId}", handleBranchStaff)
}

func handleBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == "GET" {
		resp, err := ctrl.GetBranch(ctx)
		if err != nil {
			err = enc.Encode(&schema.GetBranchResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetBranchResponse{
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
		var request *schema.AddBranchRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.AddBranch(ctx, request)
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

func handleBranchDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	vars := mux.Vars(r)
	branchId := vars["branchId"]
	if r.Method == "GET" {
		resp, err := ctrl.GetBranchDetail(ctx, branchId)
		if err != nil {
			err = enc.Encode(&schema.GetBranchDetailResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetBranchDetailResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
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
		var request *schema.UpdateBranchRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.UpdateBranch(ctx, request, cast.ToInt32(branchId))
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
	} else if r.Method == "DELETE" {
		err := ctrl.DeleteBranch(ctx, cast.ToInt32(branchId))
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

func handleBranchManager(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	vars := mux.Vars(r)
	branchId := vars["branchId"]
	if r.Method == "PUT" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		var request *schema.UpdateBranchManagerRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		err = ctrl.UpdateBranchManager(ctx, request, cast.ToInt32(branchId))
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

func handleBranchStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	vars := mux.Vars(r)
	branchId := vars["branchId"]
	if r.Method == "GET" {
		resp, err := ctrl.GetBranchStaff(ctx, branchId)
		if err != nil {
			err = enc.Encode(&schema.GetBranchStaffResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&schema.GetBranchStaffResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       resp,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
