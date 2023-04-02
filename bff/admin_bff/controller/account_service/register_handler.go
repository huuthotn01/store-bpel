package account_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/account_service"
)

var accountController IAccountBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	accountController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/account-service/account/get", handleGetAccount)
	mux.HandleFunc("/api/bff/account-service/account/add", handleAddAccount)
	mux.HandleFunc("/api/bff/account-service/account/role/update", handleUpdateRole)
	mux.HandleFunc("/api/bff/account-service/account/sign-in", handleSignIn)
	mux.HandleFunc("/api/bff/account-service/account/sign-up", handleSignUp)
}

func handleGetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&account_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Account-handleGetAccount-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		username := r.URL.Query().Get("username")
		var request = new(account_service.GetListAccountRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&account_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleGetAccount-xml.Unmarshal err %v", err),
			})
			return
		}
		request.Username = username
		account, err := accountController.GetListAccount(ctx, request)
		if err != nil {
			err = enc.Encode(&account_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleGetAccount-GetListAccount err %v", err),
			})
		} else {
			err = enc.Encode(&account_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       account,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleAddAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&account_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Account-handleAddAccount-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(account_service.AddAccountRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleAddAccount-xml.Unmarshal err %v", err),
			})
			return
		}
		err := accountController.AddAccount(ctx, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleAddAccount-AddAccount err %v", err),
			})
		} else {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleUpdateRole(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&account_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Account-handleUpdateRole-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(account_service.UpdateRoleRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleUpdateRole-xml.Unmarshal err %v", err),
			})
			return
		}
		err := accountController.UpdateRole(ctx, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleUpdateRole-AddAccount err %v", err),
			})
		} else {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&account_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Account-handleSignIn-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(account_service.SignInRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&account_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleSignIn-xml.Unmarshal err %v", err),
			})
			return
		}
		role, err := accountController.SignIn(ctx, request)
		if err != nil {
			err = enc.Encode(&account_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleSignIn-SignIn err %v", err),
			})
		} else {
			err = enc.Encode(&account_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       role,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&account_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Account-handleSignUp-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(account_service.SignUpRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleSignUp-xml.Unmarshal err %v", err),
			})
		}
		err := accountController.SignUp(ctx, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleSignUp-SignUp err %v", err),
			})
		} else {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 200,
				Message:    "OK",
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
