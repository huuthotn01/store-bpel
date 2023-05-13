package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"store-bpel/branch_service/config"
	"store-bpel/branch_service/internal/controller"
	"store-bpel/branch_service/schema"
	"store-bpel/library/kafka_lib"
)

func Consume(ctx context.Context, cfg *config.Config, ctrl controller.IBranchServiceController) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%v", cfg.KafkaHost, cfg.KafkaPort)},
		Topic:   kafka_lib.BRANCH_SERVICE_TOPIC,
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Println("Could not consume message " + err.Error())
			continue
		}

		var request *schema.AddBranchStaffRequest
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Println("Could not unmarshal value " + err.Error())
			continue
		}

		err = ctrl.AddBranchStaff(ctx, request)
		if err != nil {
			log.Println("Cannot process AddAccount" + err.Error())
			continue
		}
		log.Println("Done processing AddBranchStaff")
	}
}
