package schema

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Status string

type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`

	CartID     string `json:"card_id" bson:"card_id"`
	CustomerID string `json:"customer_id" bson:"customer_id"`

	Status Status `json:"status" bson:"status"`
	Price  int64  `json:"price" bson:"price"`

	Customer        Customer        `json:"customer" bson:"customer"`
	ShippingAddress ShippingAddress `json:"shipping_address" bson:"shipping_address"`
	BillingAddress  BillingAddress  `json:"billing_address" bson:"billing_address"`

	Items []OrderItem `json:"order_items" bson:"order_items"`
}

type OrderItem struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Price    int64  `json:"price" bson:"price"`
	Quantity int    `json:"quantity" bson:"quantity"`
}
