package db

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"order/schema"
	"order/util"
)

const ORDERCOLLECTION = "orders"

type OrderRepository struct {
	Conn *mongo.Database
}

func NewMongo() (*OrderRepository, error) {
	client, err := util.GetDBConnection()
	if err != nil {
		return nil, err
	}
	client.Connect(context.Background())
	db := client.Database("order-database")
	return &OrderRepository{
		db,
	}, nil
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *schema.Order) error {
	return nil
}

func (r *OrderRepository) ListOrderByCustomerID(ctx context.Context, id string) (*schema.Order, error) {
	return nil, nil
}
