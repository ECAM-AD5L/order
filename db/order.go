package db

import (
	"context"
	"fmt"
	"order/schema"
	"order/util"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
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

func (r *OrderRepository) ListOrderByCustomerID(ctx context.Context, id string) ([]*schema.Order, error) {
	idDoc := bson.D{{"customer_id", id}}

	orders := make([]*schema.Order, 0)
	c, err := r.Conn.Collection(ORDERCOLLECTION).Find(ctx, idDoc)

	if err != nil {
		return nil, err
	}
	for c.Next(ctx) {
		var order schema.Order
		if err = c.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (r *OrderRepository) GetOrder(ctx context.Context, id string) (*schema.Order, error) {

	objectIDS, err := primitive.ObjectIDFromHex(id)
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

func (r *OrderRepository) GetOrders(ctx context.Context) ([]*schema.Order, error) {
	orders := make([]*schema.Order, 0)
	c, err := r.Conn.Collection(ORDERCOLLECTION).Find(ctx, nil)

	if err != nil {
		return nil, err
	}
	for c.Next(ctx) {
		var order schema.Order
		if err = c.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (r *OrderRepository) UserHasItem(ctx context.Context, userId string, itemId string) (*bool, error) {
	z := new(bool)
	query := fmt.Sprintf(`{
		{$match: {"customer_id": %s}},
		{$addFields : {"orders":{$filter:{
			input: "$orders",
			as: "order",
			cond: {$eq: ["$$order.ID", %s]}
		}}}}}`, userId, itemId)

	var group interface{}
	err := bson.UnmarshalExtJSON([]byte(query), true, group)
	if err != nil {
		*z = false
		return z, err
	}
	_, err = r.Conn.Collection(ORDERCOLLECTION).Aggregate(ctx, group)
	if err != nil {
		*z = false
		return z, err
	}

	*z = false
	return z, nil
}
