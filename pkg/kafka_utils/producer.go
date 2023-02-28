package kafka

import (
	"fmt"
	"time"

	"github.com/duyledat197/go-gen-tools/config"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
)

type Producer struct {
	ClientID string
	Address  *config.ConnectionAddr
	Producer sarama.SyncProducer
}

func NewProducer(address *config.ConnectionAddr) (*Producer, error) {
	clientID := uuid.NewString()
	config := sarama.NewConfig()

	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Version = sarama.V1_0_0_0
	config.Metadata.Retry.Backoff = time.Second * 2
	config.ClientID = clientID

	producer, err := sarama.NewSyncProducer([]string{address.GetConnectionString()}, config)
	if err != nil {
		return nil, fmt.Errorf("cannot create Kafka producer: %w", err)
	}

	return &Producer{
		ClientID: clientID,
		Address:  address,
		Producer: producer,
	}, nil
}
