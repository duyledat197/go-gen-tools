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

type Subscriber struct {
	ServiceName string
	Brokers     []*config.ConnectionAddr
	consumer    *ConsumerGroup
	// Partition   int32
	Topics  []*pubsub.Topic
	Logger  *zap.Logger
	Handler func(msg *pubsub.Message, eventTime time.Time)
}

func (s *Subscriber) Init(ctx context.Context) error {
	s.consumer = &ConsumerGroup{
		serviceName: s.ServiceName,
		brokers:     s.Brokers,
		topics:      s.Topics,
		logger:      s.Logger,
		handler:     s.Handler,
	}
	return nil
}
func (s *Subscriber) Start(ctx context.Context) error {
	return s.consumer.start(ctx)
}

func (s *Subscriber) Stop(ctx context.Context) error {
	return s.consumer.stop(ctx)
}
