package account_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/schema/account_service"

	"github.com/gorilla/mux"
)

var accountController IAccountBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	accountController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/account-service/account/sign-in", handleSignIn)
	mux.HandleFunc("/api/bff/account-service/account/sign-up", handleSignUp)
	mux.HandleFunc("/api/bff/account-service/account/reset-password/create", handleCreateResetPassword)
	mux.HandleFunc("/api/bff/account-service/account/reset-password/otp", handleConfirmOTP)
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

func handleCreateResetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&account_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Account-handleCreateResetPassword-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(account_service.CreateResetPasswordRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleCreateResetPassword-xml.Unmarshal err %v", err),
			})
		}
		err := accountController.CreateResetPassword(ctx, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleCreateResetPassword-CreateResetPassword err %v", err),
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

func handleConfirmOTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&account_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Account-handleConfirmOTP-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(account_service.ConfirmOTPRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleConfirmOTP-xml.Unmarshal err %v", err),
			})
		}
		err := accountController.ConfirmOTP(ctx, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleConfirmOTP-ConfirmOTP err %v", err),
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
