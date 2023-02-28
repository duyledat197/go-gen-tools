package pubsub

import (
	"context"
)

type Subscriber interface {
	Init(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
