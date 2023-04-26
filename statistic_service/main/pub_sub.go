package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"store-bpel/library/kafka_lib"
	"store-bpel/statistic_service/internal/controller"
	"store-bpel/statistic_service/schema"
)

func Consume(ctx context.Context, ctrl controller.IStatisticServiceController) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   kafka_lib.STATISTIC_SERVICE_TOPIC,
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("Could not consume message " + err.Error())
		}

		var request *schema.AddOrderDataRequest
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			panic("Could not unmarshal value " + err.Error())
		}

		err = ctrl.AddOrderData(ctx, request)
		if err != nil {
			panic("Cannot process AddOrderData" + err.Error())
		}
		log.Println("Done processing AddOrderData")
	}
}
