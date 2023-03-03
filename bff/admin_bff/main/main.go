package main

import (
	"context"
	"encoding/xml"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/controller"
	"store-bpel/bff/admin_bff/schema/branch_service"
)

var ctrl controller.IAdminBffController
var branchAdapter adapter.IBranchServiceAdapter

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Admin BFF server started at port %d", cfg.HttpPort)

	ctrl = controller.NewController(cfg)
	branchAdapter = adapter.NewBranchAdapter(cfg)

	mux := newSOAPMux()

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), mux); err != nil {
		log.Fatal(err)
	}
	log.Printf("Admin BFF initialized successfully at port %d", cfg.HttpPort)
}

func newSOAPMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/bff/branch-service", handleBranch)
	return mux
}

func handleBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&branch_service.UpdateResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
		return
	}
	var request = new(branch_service.BranchRequest)
	if r.Method == http.MethodGet {
		bodyData := &branch_service.GetBranchRequestData{}
		request.Body = bodyData
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		}
		branch, err := branchAdapter.GetBranch(ctx, bodyData.BranchId)
		if err != nil {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		} else {
			err = enc.Encode(&branch_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       branch,
			})
		}
	} else if r.Method == http.MethodPost {
		request.Body = &branch_service.AddBranchRequestData{}
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&branch_service.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
