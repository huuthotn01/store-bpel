package main

import (
	"github.com/spf13/cast"
	"log"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	branch_controller "store-bpel/bff/admin_bff/controller/branch_service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Admin BFF server started at port %d", cfg.HttpPort)

	mux := newSOAPMux(cfg)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), mux); err != nil {
		log.Fatal(err)
	}
	log.Printf("Admin BFF initialized successfully at port %d", cfg.HttpPort)
}

func newSOAPMux(cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	branch_controller.RegisterEndpointHandler(mux, cfg)
	return mux
}
