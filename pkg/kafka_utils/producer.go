package kafka_utils

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type Producer struct {
	ClientID string
	Address  string
	Producer sarama.SyncProducer
}

func NewProducer(clientID, address string) (*Producer, error) {
	config := sarama.NewConfig()

	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Version = sarama.V1_0_0_0
	config.Metadata.Retry.Backoff = time.Second * 2
	config.ClientID = clientID

	producer, err := sarama.NewSyncProducer([]string{address}, config)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create Kafka producer")
	}
	return &Producer{
		ClientID: clientID,
		Address:  address,
		Producer: producer,
	}, nil
}
