package kafka

import (
	"context"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/pkg/pubsub"

	"github.com/Shopify/sarama"
)

type publisher struct {
	Producer *Producer
	Address  *config.ConnectionAddr
}

func (p *publisher) Publish(ctx context.Context, topic *pubsub.Topic, msg *pubsub.Message) error {
	m := &sarama.ProducerMessage{
		Topic: topic.Name,
		Value: sarama.ByteEncoder(msg.Msg),
		Key:   sarama.ByteEncoder(msg.Key),
	}
	if _, _, err := p.Producer.Producer.SendMessage(m); err != nil {
		return err
	}
	return nil
}

func (p *publisher) Connect(ctx context.Context) error {
	producer, err := NewProducer(p.Address)
	if err != nil {
		return err
	}
	p.Producer = producer
	return nil
}
func (p *publisher) Stop(ctx context.Context) error {
	return p.Producer.Producer.Close()
}

func NewPublisher(address *config.ConnectionAddr) pubsub.Publisher {
	return &publisher{
		Address: address,
	}
}
