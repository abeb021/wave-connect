package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	kp *kafka.Producer
}

func NewProducer(broker string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":broker,
	})
	if err != nil{
		return nil, err
	}

	return &Producer{kp: p}, nil
}

func (p *Producer) Close(){
	p.kp.Close()
}