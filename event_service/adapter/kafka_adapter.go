package adapter

type IKafkaAdapter interface {

}

type kafkaAdapter struct {

}

func NewKafkaAdapter() IKafkaAdapter {
	return &kafkaAdapter{}
}
