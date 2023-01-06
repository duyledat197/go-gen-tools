package kafka_utils

import (
	"time"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	ClientID  string
	Address   string
	Consumer  sarama.PartitionConsumer
	partition int32
	topic     *Topic
}

func NewConsumer(clientID, address string, partition int32, topic *Topic) (*Consumer, error) {
	cfg := sarama.NewConfig()
	cfg.ClientID = clientID

	cfg.Net.ReadTimeout = 3 * time.Second
	cfg.Consumer.Retry.Backoff = 200 * time.Millisecond
	cfg.Consumer.Return.Errors = true
	cfg.Metadata.Retry.Max = 5

	master, err := sarama.NewConsumer([]string{address}, cfg)
	if err != nil {
		return nil, err
	}
	consumer, err := master.ConsumePartition(topic.Name, partition, sarama.OffsetNewest)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		ClientID:  clientID,
		Address:   address,
		topic:     topic,
		partition: partition,
		Consumer:  consumer,
	}, nil
}
