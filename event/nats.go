package event

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"order/schema"
	"os"

	"github.com/nats-io/go-nats"
)

type NatsEventStore struct {
	nc                       *nats.Conn
	orderCreatedSubscription *nats.Subscription
	orderCreatedChan         chan OrderCreatedMessage
}

func NewNats() (*NatsEventStore, error) {
	nc, err := nats.Connect(fmt.Sprintf("nats://%s", os.Getenv("NATS_HOST")))
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

func (e *NatsEventStore) Close() {
	if e.nc != nil {
		e.nc.Close()
	}
	if e.orderCreatedSubscription != nil {
		e.orderCreatedSubscription.Unsubscribe()
	}
	close(e.orderCreatedChan)
}

func (e *NatsEventStore) SubscribeOrderCreated() (<-chan OrderCreatedMessage, error) {
	m := OrderCreatedMessage{}
	e.orderCreatedChan = make(chan OrderCreatedMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	e.orderCreatedSubscription, err = e.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}
	// Decode message
	go func() {
		for {
			select {
			case msg := <-ch:
				e.readMessage(msg.Data, &m)
				e.orderCreatedChan <- m
			}
		}
	}()
	return (<-chan OrderCreatedMessage)(e.orderCreatedChan), nil
}

func (e *NatsEventStore) OnOrderCreated(f func(OrderCreatedMessage)) (err error) {
	m := OrderCreatedMessage{}
	e.orderCreatedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		e.readMessage(msg.Data, &m)
		f(m)
	})
	return
}

func (e *NatsEventStore) PublishOrderCreated(order schema.Order) error {
	m := OrderCreatedMessage{order.ID.Hex(), order.CreatedAt}
	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (e *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
