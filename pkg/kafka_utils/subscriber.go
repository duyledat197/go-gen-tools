package kafka

import (
	"context"
	"time"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/pkg/pubsub"

	"go.uber.org/zap"
)

type consumerType int

const (
	ConsumerType_GROUP consumerType = iota
	ConsumerType_PARTITION
)

type subscriber struct {
	ServiceName string
	Brokers     []*config.ConnectionAddr
	consumer    *ConsumerGroup
	// Partition   int32
	Topics  []*pubsub.Topic
	Logger  *zap.Logger
	Handler func(msg *pubsub.Message, eventTime time.Time)
}

func (s *subscriber) Connect(ctx context.Context) error {
	s.consumer = &ConsumerGroup{
		serviceName: s.ServiceName,
		brokers:     s.Brokers,
		topics:      s.Topics,
		logger:      s.Logger,
		handler:     s.Handler,
	}

	return s.consumer.Connect(ctx)
}

func (s *subscriber) Close(ctx context.Context) error {
	return s.consumer.Close(ctx)
}
