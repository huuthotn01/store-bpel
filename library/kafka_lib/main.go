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
	Consume(ctx context.Context, topic string) ([]byte, error)
}

type kafkaLib struct {
	host string
}

func NewKafkaLib() IKafkaLib {
	return &kafkaLib{
		host: "localhost",
	}
}

type TestMessage struct {
	Counter int
	IsOkay  string
}

func (k *kafkaLib) Publish(ctx context.Context, topic string, msg []byte) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
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

func (k *kafkaLib) Consume(ctx context.Context, topic string) ([]byte, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("localhost:9092")},
		Topic:   topic,
		GroupID: "group-1",
	})

	msg, err := r.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}
	return msg.Value, nil
}
