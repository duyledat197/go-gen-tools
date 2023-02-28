package kafka

import (
	"context"
	"fmt"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/pkg/pubsub"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
)

type producer struct {
	clientID string
	address  *config.ConnectionAddr
	producer sarama.SyncProducer
}

func (p *producer) connect(ctx context.Context) error {
	clientID := uuid.NewString()
	config := sarama.NewConfig()
	config.ClientID = clientID

	producer, err := sarama.NewSyncProducer([]string{p.address.GetConnectionString()}, config)
	if err != nil {
		return fmt.Errorf("cannot create Kafka producer: %w", err)
	}
	p.producer = producer
	p.clientID = clientID
	return nil
}
func (p *producer) stop(ctx context.Context) error {
	return p.producer.Close()
}

func (p *producer) publish(ctx context.Context, topic *pubsub.Topic, msg *pubsub.Message) error {
	m := &sarama.ProducerMessage{
		Topic: topic.Name,
		Value: sarama.ByteEncoder(msg.Msg),
		Key:   sarama.ByteEncoder(msg.Key),
	}
	if _, _, err := p.producer.SendMessage(m); err != nil {
		return fmt.Errorf("p.producer.SendMessage: %w", err)
	}
	return nil
}
