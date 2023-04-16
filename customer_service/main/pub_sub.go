package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"store-bpel/customer_service/controller"
	"store-bpel/customer_service/schema"
	"store-bpel/library/kafka_lib"
)

func Consume(ctx context.Context, ctrl controller.ICustomerServiceController) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   kafka_lib.CUSTOMER_SERVICE_TOPIC,
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("Could not consume message " + err.Error())
		}

		var request *schema.AddCustomerRequest
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			panic("Could not unmarshal value " + err.Error())
		}
		log.Println(request)

		err = ctrl.AddCustomer(ctx, request)
		if err != nil {
			panic("Cannot process AddCustomer" + err.Error())
		}
		log.Println("Done processing AddCustomer")
	}
}
