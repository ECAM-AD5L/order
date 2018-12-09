package db

import (
	"context"
	"order/schema"
	"order/util"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
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
	_, err := r.Conn.Collection(ORDERCOLLECTION).InsertOne(
		ctx,
		order)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) ListOrderByCustomerID(ctx context.Context, id string) (*schema.Order, error) {
	return nil, nil
}

func (r *OrderRepository) GetOrder(ctx context.Context, id string) (*schema.Order, error) {

	objectIDS, err := objectid.FromHex(id)
	if err != nil {
		return nil, err
	}
	idDoc := bson.D{{"_id", objectIDS}}
	var order schema.Order
	err = r.Conn.Collection(ORDERCOLLECTION).FindOne(ctx, idDoc).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
