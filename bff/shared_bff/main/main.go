package main

import (
	"github.com/spf13/cast"
	"log"
	"net/http"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/controller"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Shared BFF server started at port %d", cfg.HttpPort)

	mux := newSOAPMux(cfg)

	if err = http.ListenAndServe(":"+cast.ToString(cfg.HttpPort), mux); err != nil {
		log.Fatal(err)
	}
	log.Printf("Shared BFF initialized successfully at port %d", cfg.HttpPort)
}

func newSOAPMux(cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	controller.RegisterEndpointHandler(mux, cfg)
	return mux
}
