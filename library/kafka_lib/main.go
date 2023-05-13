package kafka_lib

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type IKafkaLib interface {
	Publish(ctx context.Context, topic string, msg []byte) error
}

type kafkaLib struct {
	host string
	port int
}

func NewKafkaLib(host string, port int) IKafkaLib {
	return &kafkaLib{
		host: host,
		port: port,
	}
}

func (k *kafkaLib) Publish(ctx context.Context, topic string, msg []byte) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{fmt.Sprintf("%s:%v", k.host, k.port)},
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
