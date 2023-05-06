package kafka_lib

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	BROKER_HOST = "broker"
	BROKER_PORT = "29092"
)

type IKafkaLib interface {
	Publish(ctx context.Context, topic string, msg []byte) error
}

type kafkaLib struct {
	host string
}

func NewKafkaLib() IKafkaLib {
	return &kafkaLib{
		host: BROKER_HOST,
	}
}

func (k *kafkaLib) Publish(ctx context.Context, topic string, msg []byte) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", BROKER_HOST, BROKER_PORT)},
		Topic:   topic,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(time.Now().String()),
		Value: msg,
	})
	if err != nil {
		return err
	}
	log.Printf("Published to topic %s", topic)
	return nil
}
