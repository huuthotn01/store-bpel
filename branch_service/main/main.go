package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/branch_service/config"
	"store-bpel/branch_service/internal/controller"
	"store-bpel/branch_service/schema"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
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

	ctx := context.Background()
	go Consume(ctx, cfg, ctrl)

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
	r.HandleFunc("/api/branch-service/image/{branchId}", handleBranchImage)
}

func handleBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	if r.Method == http.MethodGet {
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
	} else if r.Method == http.MethodPost {
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
	if r.Method == http.MethodGet {
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
	} else if r.Method == http.MethodPut {
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
		err = ctrl.UpdateBranch(ctx, request, branchId)
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
		err := ctrl.DeleteBranch(ctx, branchId)
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
	if r.Method == http.MethodPut {
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
		err = ctrl.UpdateBranchManager(ctx, request, branchId)
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
	if r.Method == http.MethodGet {
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

// TODO IMPROVE NOT USED YET
func handleBranchImage(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()
	w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	// vars := mux.Vars(r)
	// branchId := vars["branchId"]
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20) // max size 10MB
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}

		file, handler, err := r.FormFile("images")
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		defer file.Close()

		log.Printf("Filename: %s", handler.Filename)
		log.Printf("File size: %d", handler.Size)
		log.Printf("File header: %v", handler.Header)

		tempFile, err := ioutil.TempFile("./", "uploaded-*.png")
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}
		_, err = tempFile.Write(fileBytes)
		if err != nil {
			err = enc.Encode(&schema.UpdateResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
			return
		}

		log.Println("Uploaded file successfully")

		err = enc.Encode(&schema.UpdateResponse{
			StatusCode: 200,
			Message:    "OK",
		})
		//err := ctrl.UploadBranchImage(ctx, branchId)
		//if err != nil {
		//	err = enc.Encode(&schema.UpdateResponse{
		//		StatusCode: 500,
		//		Message:    err.Error(),
		//	})
		//} else {
		//	err = enc.Encode(&schema.UpdateResponse{
		//		StatusCode: 200,
		//		Message:    "OK",
		//	})
		//}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
