package kafka_utils

import (
	"go.uber.org/zap"
)

type Subscriber interface {
	Subscribe(fn func(msg []byte))
}

type subscriber struct {
	Address   string
	Consumer  *Consumer
	Partition int32
	Topic     *Topic
	logger    *zap.Logger
}

func (s *subscriber) Subscribe(fn func(msg []byte)) {
	for {
		select {
		case msg := <-s.Consumer.Consumer.Messages():
			s.logger.Info("Received messages", zap.String(string(msg.Key), string(msg.Value)))
			fn(msg.Value)
		case err := <-s.Consumer.Consumer.Errors():
			s.logger.Error("consumer got error", zap.Error(err))
		}
	}
}

func NewSubscriber(topic *Topic, clientID, address string, partition int32) (Subscriber, error) {

	consumer, err := NewConsumer(clientID, address, partition, topic)
	if err != nil {
		return nil, err
	}

	return &subscriber{
		Consumer:  consumer,
		Address:   address,
		Partition: partition,
		Topic:     topic,
	}, nil
}
