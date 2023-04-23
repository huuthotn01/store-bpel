package customer_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"store-bpel/bff/admin_bff/common"
	"store-bpel/bff/admin_bff/schema/event_service"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/schema/customer_service"
	"time"
)

var customerController ICustomerBffController

func RegisterEndpointHandler(r *mux.Router, cfg *config.Config) {
	// init controller
	customerController = NewController(cfg)
	// register handler
	r.HandleFunc("/api/bff/customer-service/customer/get", handleGetCustomer)
	r.HandleFunc("/api/bff/customer-service/customer/add", handleAddCustomer)
	r.HandleFunc("/api/bff/customer-service/customer/update", handleUpdateCustomer)
	r.HandleFunc("/api/bff/customer-service/image:upload", handleUploadImage)
	r.HandleFunc("/api/bff/customer-service/image:delete", handleDeleteImage)
}

func handleGetCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&customer_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Customer-handleGetCustomer-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(customer_service.GetCustomerInfoRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&customer_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleGetCustomer-xml.Unmarshal err %v", err),
			})
			return
		}
		customer, err := customerController.GetCustomer(ctx, request)
		if err != nil {
			err = enc.Encode(&customer_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleGetCustomer-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&customer_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       customer,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAddCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&customer_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Customer-handleAddCustomer-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(customer_service.AddCustomerRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleAddCustomer-xml.Unmarshal err %v", err),
			})
			return
		}
		err = customerController.AddCustomer(ctx, request)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleAddCustomer-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&customer_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Customer-handleUpdateCustomer-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(customer_service.UpdateCustomerInfoRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleUpdateCustomer-xml.Unmarshal err %v", err),
			})
			return
		}
		err = customerController.UpdateCustomer(ctx, request)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleUpdateCustomer-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	if r.Method == http.MethodPost {
		var (
			username = r.FormValue("username")
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
		err = os.MkdirAll(fmt.Sprintf("../uploads/%s", username), os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a new file in the uploads directory
		relativePath := fmt.Sprintf("uploads/%s/%d%s", username, time.Now().Unix(), filepath.Ext(fileHeader.Filename))
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

		imgPath := "/store-bpel/bff/customer_bff/" + relativePath
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = customerController.UploadImage(ctx, &customer_service.UploadImageRequest{
			Username: username,
			ImageUrl: imgPath,
		})
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleUploadImage-UploadImage err %v", err),
			})
		} else {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
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
		err = enc.Encode(&customer_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Customer-handleDeleteImage-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(customer_service.DeleteImageRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleDeleteImage-xml.Unmarshal err %v", err),
			})
			return
		}
		err = customerController.DeleteImage(ctx, request.Username)
		if err != nil {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Customer-handleDeleteImage-GetCustomer err %v", err),
			})
		} else {
			err = enc.Encode(&customer_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
