package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/pkg/pubsub"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type ConsumerGroup struct {
	serviceName string
	brokers     []*config.ConnectionAddr
	topics      []*pubsub.Topic
	client      sarama.ConsumerGroup
	logger      *zap.Logger
	handler     func(msg *pubsub.Message, eventTime time.Time)
}

func (g *ConsumerGroup) start(ctx context.Context) error {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	var addrs []string
	for _, broker := range g.brokers {
		addrs = append(addrs, broker.GetConnectionString())
	}

	client, err := sarama.NewConsumerGroup(addrs, g.serviceName, config)
	if err != nil {
		return fmt.Errorf("error creating consumer group client: %w", err)
	}

	g.client = client
	g.subscribe()
	return nil
}

func (g *ConsumerGroup) stop(ctx context.Context) error {
	return g.client.Close()
}

func (g *ConsumerGroup) subscribe() {
	ctx := context.Background()
	consumer := consumerGroupHandler{
		ready:  make(chan bool),
		fn:     g.handler,
		logger: g.logger,
	}
	var topics []string
	for _, topic := range g.topics {
		topics = append(topics, topic.Name)
	}
	go func() {
		for {
			// TODO: `Consume` should be called inside an infinite loop, when a
			// TODO: server-side rebalance happens, the consumer session will need to be
			// TODO: recreated to get the new claims
			if err := g.client.Consume(ctx, topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
	<-consumer.ready
}

type consumerGroupHandler struct {
	ready  chan bool
	fn     func(msg *pubsub.Message, eventTime time.Time)
	logger *zap.Logger
}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	h.logger.Info("Sarama consumer up and running!...")
	close(h.ready)
	return nil
}

func (h *consumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// TODO: ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		h.logger.Sugar().Infof("Message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
		h.fn(&pubsub.Message{
			Key: message.Key,
			Msg: message.Value,
		}, message.Timestamp)
	}
	return nil
}
