package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"store-bpel/cart_service/internal/controller"
	"store-bpel/cart_service/schema"
	"store-bpel/library/kafka_lib"
)

func Consume(ctx context.Context, ctrl controller.ICartServiceController) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   kafka_lib.CART_SERVICE_TOPIC,
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("Could not consume message " + err.Error())
		}

		var request *schema.AddCartRequest
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			panic("Could not unmarshal value " + err.Error())
		}

		err = ctrl.AddCart(ctx, request.CustomerId)
		if err != nil {
			panic("Cannot process AddCart" + err.Error())
		}
		log.Println("Done processing AddCart")
	}
}
