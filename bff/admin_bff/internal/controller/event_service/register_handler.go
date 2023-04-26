package event_service

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
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/internal/common"
	"store-bpel/bff/admin_bff/schema/event_service"
	"time"
)

var eventController IEventBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	eventController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/event-service/add-event", handleAddEvent)
	mux.HandleFunc("/api/bff/event-service/update-event", handleUpdateEvent)
	mux.HandleFunc("/api/bff/event-service/delete-event", handleDeleteEvent)
	mux.HandleFunc("/api/bff/event-service/image:upload", handleUploadImage)
	mux.HandleFunc("/api/bff/event-service/image:delete", handleDeleteImage)
}

func handleAddEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	if r.Method == http.MethodPost {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleAddEvent-ioutil.ReadAll err %v", err),
			})
		}

		var req *event_service.AddEventRequest
		err = xml.Unmarshal(payload, &req)

		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleAddEvent-json.Unmarshal err %v", err),
			})
		}
		err = eventController.AddEvent(ctx, req)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleAddEvent-AddEvent err %v", err),
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

func handleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	if r.Method == http.MethodPost {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleUpdateEvent-ioutil.ReadAll err %v", err),
			})
		}

		var req *event_service.UpdateEventRequest
		err = xml.Unmarshal(payload, &req)

		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleUpdateEvent-json.Unmarshal err %v", err),
			})
		}
		err = eventController.UpdateEvent(ctx, req)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleUpdateEvent-UpdateEvent err %v", err),
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

func handleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	if r.Method == http.MethodPost {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleDeleteEvent-ioutil.ReadAll err %v", err),
			})
		}

		var req *event_service.DeleteEventRequest
		err = xml.Unmarshal(payload, &req)

		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleDeleteEvent-json.Unmarshal err %v", err),
			})
		}
		err = eventController.DeleteEvent(ctx, req)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleDeleteEvent-DeleteEvent err %v", err),
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

func handleUploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)

	if r.Method == http.MethodPost {
		var (
			eventId = r.FormValue("eventId")
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
		err = os.MkdirAll(fmt.Sprintf("../uploads/%s", eventId), os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a new file in the uploads directory
		relativePath := fmt.Sprintf("uploads/%s/%d%s", eventId, time.Now().Unix(), filepath.Ext(fileHeader.Filename))
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

		err = eventController.UploadImage(ctx, &event_service.UploadImageRequest{
			EventId:  eventId,
			ImageUrl: imgPath,
		})
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleUploadImage-UploadImage err %v", err),
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

	if r.Method == http.MethodPost {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleDeleteImage-ioutil.ReadAll err %v", err),
			})
		}
		var req *event_service.DeleteImageRequest
		err = xml.Unmarshal(payload, &req)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleDeleteImage-json.Unmarshal err %v", err),
			})
		}
		err = eventController.DeleteImage(ctx, req)
		if err != nil {
			err = enc.Encode(&event_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Event-handleDeleteImage-DeleteEvent err %v", err),
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
