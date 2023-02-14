package adapter

import (
	"context"
	"net/http"
	"store-bpel/goods_service/config"
)

type IKafkaAdapter interface {
	Publish(ctx context.Context, topic string) error
}

type kafkaAdapter struct {
	httpClient *http.Client
	port int
}

func NewKafkaAdapter(cfg *config.Config) IKafkaAdapter {
	return &kafkaAdapter{
		httpClient: &http.Client{},
		port: cfg.KafkaPort,
	}
}

func (a *kafkaAdapter) Publish(ctx context.Context, topic string) error {
	return nil
}
