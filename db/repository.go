package db

import (
	"context"
	"order/schema"
)

type Repository interface {
	//OK
	CreateOrder(ctx context.Context, order *schema.Order) error

	ListOrderByCustomerID(ctx context.Context, id string) ([]*schema.Order, error)

	//OK
	GetOrder(ctx context.Context, id string) (*schema.Order, error)

	//OK
	GetOrders(ctx context.Context) ([]*schema.Order, error)

	UserHasItem(ctx context.Context, userId string, itemId string) (*bool, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func CreateOrder(ctx context.Context, order *schema.Order) error {
	return impl.CreateOrder(ctx, order)
}

func ListOrderByCustomerID(ctx context.Context, id string) ([]*schema.Order, error) {
	return impl.ListOrderByCustomerID(ctx, id)
}

func GetOrder(ctx context.Context, id string) (*schema.Order, error) {
	return impl.GetOrder(ctx, id)
}

func GetOrders(ctx context.Context) ([]*schema.Order, error) {
	return impl.GetOrders(ctx)
}

func UserHasItem(ctx context.Context, userId string, itemId string) (*bool, error) {
	return impl.UserHasItem(ctx, userId, itemId)
}
