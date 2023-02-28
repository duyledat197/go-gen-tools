package pubsub

import "context"

type Publisher interface {
	Publish(ctx context.Context, msg *Message) error
	Connect(ctx context.Context) error
	Stop(ctx context.Context) error
}
