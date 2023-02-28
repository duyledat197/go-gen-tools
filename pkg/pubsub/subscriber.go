package pubsub

type Subscriber interface {
	Subscribe(fn func(msg []byte))
}
