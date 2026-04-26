package kafka

import (
	"errors"

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

func (p *Producer) Send(topic, key string, value []byte) error{
	deliveryChan := make(chan kafka.Event)

	err := p.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
			Partition: kafka.PartitionAny,
		},
		Key: []byte(key),
		Value: value,
	}, deliveryChan)

	if err != nil{
		return err
	}

	e := <-deliveryChan
	m, ok := e.(*kafka.Message)
	if !ok{
		return errors.New("unexpected error")
	}

	if m.TopicPartition.Error != nil{
		return m.TopicPartition.Error
	}

	return nil
}