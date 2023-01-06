package kafka_utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

type ConsumerGroupHandler interface {
	Subscribe(fn func(msg []byte, timeEvent time.Time, key []byte))
}

type consumerGroup struct {
	topics string
	client sarama.ConsumerGroup
}

func NewConsumerGroup(assignor, brokers, groupId, topics string, oldest bool) (ConsumerGroupHandler, error) {
	config := sarama.NewConfig()

	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		fmt.Printf("Unrecognized consumer group partition assignor: %s\n", assignor)
	}

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), groupId, config)
	if err != nil {
		fmt.Printf("Error creating consumer group client: %v\n", err)
		return nil, err
	}

	return &consumerGroup{
		topics: topics,
		client: client,
	}, nil
}

func (s *consumerGroup) Subscribe(fn func(msg []byte, timeEvent time.Time, key []byte)) {
	ctx, cancel := context.WithCancel(context.Background())
	consumer := ConsumerGroup{
		ready: make(chan bool),
		fn:    fn,
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := s.client.Consume(ctx, strings.Split(s.topics, ","), &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err := s.client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

// Consumer represents a Sarama consumer group consumer
type ConsumerGroup struct {
	ready chan bool
	fn    func(msg []byte, timeEvent time.Time, key []byte)
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *ConsumerGroup) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *ConsumerGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *ConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
		consumer.fn(message.Value, message.Timestamp, message.Key)
	}
	return nil
}
