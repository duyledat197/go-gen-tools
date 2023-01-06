package kafka_utils

import (
	"github.com/Shopify/sarama"
)

type Publisher interface {
	Publish(topic *Topic, key []byte, msg []byte) error
}

type publisher struct {
	Producer *Producer
	Address  string
	ClientID string
}

func (p *publisher) Publish(topic *Topic, key []byte, msg []byte) error {
	m := &sarama.ProducerMessage{
		Topic: topic.Name,
		Value: sarama.ByteEncoder(msg),
		Key:   sarama.ByteEncoder(key),
	}
	if _, _, err := p.Producer.Producer.SendMessage(m); err != nil {
		return err
	}
	return nil
}

func NewPublisher(clientID, address string) (Publisher, error) {
	producer, err := NewProducer(clientID, address)
	if err != nil {
		return nil, err
	}
	return &publisher{
		Address:  address,
		ClientID: clientID,
		Producer: producer,
	}, nil
}
