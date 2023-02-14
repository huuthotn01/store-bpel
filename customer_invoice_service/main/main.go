package main

import (
	"github.com/spf13/cast"
	"log"
	"net/http"
	"store-bpel/customer_invoice_service/config"
	"store-bpel/customer_invoice_service/controller"
)

var ctrl controller.ICustomerInvoiceServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Customer Service server started at port %d", cfg.HttpPort)

	db, err := DbConnect(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ctrl = controller.NewController(cfg, db)
	registerEndpoint()

	if err = http.ListenAndServe(":" + cast.ToString(cfg.HttpPort), nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("Customer Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint() {
	// http.HandleFunc({api}, {handleFunc})
}
