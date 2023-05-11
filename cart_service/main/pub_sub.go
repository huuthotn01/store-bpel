package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"store-bpel/cart_service/internal/controller"
	"store-bpel/cart_service/schema"
	"store-bpel/library/kafka_lib"
)

func Consume(ctx context.Context, ctrl controller.ICartServiceController) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", kafka_lib.BROKER_HOST, kafka_lib.BROKER_PORT)},
		Topic:   kafka_lib.CART_SERVICE_TOPIC,
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Println("Could not consume message " + err.Error())
			continue
		}

		var request *schema.AddCartRequest
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Println("Could not unmarshal value " + err.Error())
			continue
		}

		err = ctrl.AddCart(ctx, request.CustomerId)
		if err != nil {
			log.Println("Cannot process AddCart" + err.Error())
			continue
		}
		log.Println("Done processing AddCart")
	}
}
