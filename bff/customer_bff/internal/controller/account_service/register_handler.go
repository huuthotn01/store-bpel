package account_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/schema/account_service"

	"github.com/gorilla/mux"
)

var accountController IAccountBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	accountController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/account-service/account/change-password", handleChangePassword)
}

func handleChangePassword(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&account_service.UpdateResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Account-handleChangePassword-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(account_service.ChangePasswordRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleChangePassword-xml.Unmarshal err %v", err),
			})
			return
		}
		err := accountController.ChangePassword(ctx, request)
		if err != nil {
			err = enc.Encode(&account_service.UpdateResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Account-handleChangePassword-ChangePassword err %v", err),
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
