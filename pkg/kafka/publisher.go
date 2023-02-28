package kafka

import (
	"context"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/pkg/pubsub"
)

type Publisher struct {
	producer *producer
	Address  *config.ConnectionAddr
	Topic    *pubsub.Topic
}

func (p *Publisher) Publish(ctx context.Context, msg *pubsub.Message) error {
	return p.producer.publish(ctx, p.Topic, msg)
}

func (p *Publisher) Connect(ctx context.Context) error {
	p.producer = &producer{
		address: p.Address,
	}
	return p.producer.connect(ctx)
}

func (p *Publisher) Stop(ctx context.Context) error {
	return p.producer.stop(ctx)
}
