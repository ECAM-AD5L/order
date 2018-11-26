package event

import "order/schema"

type EventStore interface {
	Close()
	PublishOrderCreated(order schema.Order) error
	SubscribeOrderCreated() (<-chan OrderCreatedMessage, error)
	OnOrderCreated(f func(OrderCreatedMessage)) error
}

var impl EventStore

func SetEventStore(es EventStore) {
	impl = es
}

func Close() {
	impl.Close()
}

func PublishOrderCreated(order schema.Order) error {
	return impl.PublishOrderCreated(order)
}

func SubscribeOrderCreated() (<-chan OrderCreatedMessage, error) {
	return impl.SubscribeOrderCreated()
}

func OnOrderCreated(f func(message OrderCreatedMessage)) error {
	return impl.OnOrderCreated(f)
}
