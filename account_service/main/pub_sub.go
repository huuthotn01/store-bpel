package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"store-bpel/account_service/config"
	"store-bpel/account_service/internal/controller"
	"store-bpel/account_service/schema"
	"store-bpel/library/kafka_lib"
)

func Consume(ctx context.Context, cfg *config.Config, ctrl controller.IAccountServiceController) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%v", cfg.KafkaHost, cfg.KafkaPort)},
		Topic:   kafka_lib.ACCOUNT_SERVICE_TOPIC,
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Printf("Could not consume message " + err.Error())
			continue
		}

		var request *schema.AddAccountRequest
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Printf("Could not unmarshal value " + err.Error())
			continue
		}

		err = ctrl.AddAccount(ctx, request)
		if err != nil {
			log.Printf("Cannot process AddAccount" + err.Error())
			continue
		}
		log.Println("Done processing AddAccount")
	}
}
