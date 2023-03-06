package kafka

import (
	"context"
	"fmt"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/pkg/pubsub"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
)

type consumer struct {
	clientID  string
	brokers   []*config.ConnectionAddr
	client    sarama.PartitionConsumer
	Partition int32
	Topic     *pubsub.Topic
}

func (c *consumer) connect(ctx context.Context) error {
	cfg := sarama.NewConfig()
	cfg.ClientID = uuid.NewString()
	var addrs []string
	for _, broker := range c.brokers {
		addrs = append(addrs, broker.GetConnectionString())
	}
	client, err := sarama.NewConsumer(addrs, cfg)
	if err != nil {
		return fmt.Errorf("sarama.NewConsumer: %w", err)
	}
	consumer, err := client.ConsumePartition(c.Topic.Name, c.Partition, sarama.OffsetNewest)
	if err != nil {
		return fmt.Errorf("client.ConsumePartition: %w", err)
	}
	c.client = consumer
	return nil
}

func (g *consumer) stop(ctx context.Context) error {
	return g.client.Close()
}
