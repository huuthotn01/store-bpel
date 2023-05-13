package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"store-bpel/library/kafka_lib"
	"store-bpel/statistic_service/config"
	"store-bpel/statistic_service/internal/controller"
	"store-bpel/statistic_service/schema"
)

func Consume(ctx context.Context, cfg *config.Config, ctrl controller.IStatisticServiceController) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%v", cfg.KafkaHost, cfg.KafkaPort)},
		Topic:   kafka_lib.STATISTIC_SERVICE_TOPIC,
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Println("Could not consume message " + err.Error())
			continue
		}

		var request *schema.AddOrderDataRequest
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Println("Could not unmarshal value " + err.Error())
			continue
		}

		err = ctrl.AddOrderData(ctx, request)
		if err != nil {
			log.Println("Cannot process AddOrderData" + err.Error())
			continue
		}
		log.Println("Done processing AddOrderData")
	}
}
