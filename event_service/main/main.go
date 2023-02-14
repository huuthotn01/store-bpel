package main

import (
	"github.com/spf13/cast"
	"log"
	"net/http"
	"store-bpel/event_service/config"
	"store-bpel/event_service/controller"
)

var ctrl controller.IEventServiceController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Event Service server started at port %d", cfg.HttpPort)

	db, err := DbConnect(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ctrl = controller.NewController(cfg, db)
	registerEndpoint()

	if err = http.ListenAndServe(":" + cast.ToString(cfg.HttpPort), nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("Event Service initialized successfully at port %d", cfg.HttpPort)
}

func registerEndpoint() {
	// http.HandleFunc({api}, {handleFunc})
}
