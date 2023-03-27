package staff_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/staff_service"
)

var staffController IStaffBffController

func RegisterEndpointHandler(mux *http.ServeMux, cfg *config.Config) {
	// init controller
	staffController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/staff-service/get-staff", handleGetStaff)
	mux.HandleFunc("/api/bff/staff-service/get-staff-detail", handleGetStaffDetail)
	mux.HandleFunc("/api/bff/staff-service/add-staff", handleAddStaff)
	mux.HandleFunc("/api/bff/staff-service/update-staff", handleUpdateStaff)
	mux.HandleFunc("/api/bff/staff-service/delete-staff", handleDeleteStaff)
	mux.HandleFunc("/api/bff/staff-service/get-staff-attendance", handleGetStaffAttendance)
	mux.HandleFunc("/api/bff/staff-service/create-add-request", handleCreateAddRequest)
	mux.HandleFunc("/api/bff/staff-service/create-delete-request", handleCreateDeleteRequest)
	mux.HandleFunc("/api/bff/staff-service/update-request-status", handleUpdateRequestStatus)
	mux.HandleFunc("/api/bff/staff-service/get-request-list", handleGetRequestList)
}

func handleGetStaff(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&staff_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Staff-handleGetStaff-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(staff_service.GetStaffRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleGetStaff-xml.Unmarshal err %v", err),
			})
		}
		staff, err := staffController.GetStaff(ctx, request)
		if err != nil {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleGetStaff-GetStaff err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       staff,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetStaffDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&staff_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Staff-handleGetStaffDetail-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(staff_service.GetStaffRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleGetStaffDetail-xml.Unmarshal err %v", err),
			})
		}
		staff, err := staffController.GetStaffDetail(ctx, request)
		if err != nil {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleGetStaffDetail-GetStaff err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       staff,
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
		err = enc.Encode(&staff_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Staff-handleAddStaff-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(staff_service.AddStaffRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleAddStaff-xml.Unmarshal err %v", err),
			})
		}
		err = staffController.AddStaff(ctx, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleAddStaff-AddStaff err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.UpdateResponse{
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
		err = enc.Encode(&staff_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Staff-handleUpdateStaff-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(staff_service.UpdateStaffRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleUpdateStaff-xml.Unmarshal err %v", err),
			})
		}
		err = staffController.UpdateStaff(ctx, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleUpdateStaff-AddStaff err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.UpdateResponse{
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
		err = enc.Encode(&staff_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Staff-handleDeleteStaff-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(staff_service.CreateDeleteRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleDeleteStaff-xml.Unmarshal err %v", err),
			})
		}
		err = staffController.DeleteStaff(ctx, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleDeleteStaff-CreateDeleteRequest err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetStaffAttendance(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&staff_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Staff-handleGetStaffAttendance-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(staff_service.GetStaffAttendanceRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleGetStaffAttendance-xml.Unmarshal err %v", err),
			})
		}
		attendance, err := staffController.GetStaffAttendance(ctx, request)
		if err != nil {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleGetStaffAttendance-GetStaffAttendance err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       attendance,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleCreateAddRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&staff_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Staff-handleCreateAddRequest-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(staff_service.CreateAddRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleCreateAddRequest-xml.Unmarshal err %v", err),
			})
		}
		err = staffController.CreateAddRequest(ctx, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleCreateAddRequest-CreateAddRequest err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleCreateDeleteRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&staff_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Staff-handleCreateDeleteRequest-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(staff_service.CreateDeleteRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleCreateDeleteRequest-xml.Unmarshal err %v", err),
			})
		}
		err = staffController.CreateDeleteRequest(ctx, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleCreateDeleteRequest-CreateDeleteRequest err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateRequestStatus(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&staff_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Staff-handleUpdateRequestStatus-ioutil.ReadAll err %v", err),
		})
		return
	}

	if r.Method == http.MethodPost {
		var request = new(staff_service.UpdateRequestStatusRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleUpdateRequestStatus-xml.Unmarshal err %v", err),
			})
		}
		err = staffController.UpdateRequestStatus(ctx, request)
		if err != nil {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleUpdateRequestStatus-AddBranch err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetRequestList(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	if r.Method == http.MethodPost {
		requests, err := staffController.GetRequestList(ctx)
		if err != nil {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Staff-handleGetRequestList-GetRequestList err %v", err),
			})
		} else {
			err = enc.Encode(&staff_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       requests,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
