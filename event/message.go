package event

import "time"

type Message interface {
	Key() string
}

type OrderCreatedMessage struct {
	ID        string
	CreatedAt time.Time
}

func (m *OrderCreatedMessage) Key() string {
	return "order.created"
}
