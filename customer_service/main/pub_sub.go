package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"store-bpel/customer_service/internal/controller"
	"store-bpel/customer_service/schema"
	"store-bpel/library/kafka_lib"
)

func Consume(ctx context.Context, ctrl controller.ICustomerServiceController) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", kafka_lib.BROKER_HOST, kafka_lib.BROKER_PORT)},
		Topic:   kafka_lib.CUSTOMER_SERVICE_TOPIC,
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Println("Could not consume message " + err.Error())
		}

		var request *schema.AddCustomerRequest
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Println("Could not unmarshal value " + err.Error())
		}

		err = ctrl.AddCustomer(ctx, request)
		if err != nil {
			log.Println("Cannot process AddCustomer" + err.Error())
		}
		log.Println("Done processing AddCustomer")
	}
}
