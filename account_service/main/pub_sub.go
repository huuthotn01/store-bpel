package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"store-bpel/library/kafka_lib"
)

type TestMessage struct {
	Counter int
	IsOkay  string
}

func Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   kafka_lib.ACCOUNT_SERVICE_TOPIC,
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("Could not consume message " + err.Error())
		}

		fmt.Println(msg.Value)
		var valStr *TestMessage
		err = json.Unmarshal(msg.Value, &valStr)
		if err != nil {
			panic("Could not unmarshal value " + err.Error())
		}
		fmt.Println(valStr)
	}
}
