package db

import (
	"context"
	"order/schema"
)

type Repository interface {
	CreateOrder(ctx context.Context, order *schema.Order) error
	ListOrderByCustomerID(ctx context.Context, id string) (*schema.Order, error)
	GetOrder(ctx context.Context, id string) (*schema.Order, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}
func CreateOrder(ctx context.Context, order *schema.Order) error {
	return impl.CreateOrder(ctx, order)
}

func ListOrderByCustomerID(ctx context.Context, id string) (*schema.Order, error) {
	return impl.ListOrderByCustomerID(ctx, id)
}

func GetOrder(ctx context.Context, id string) (*schema.Order, error) {
	return impl.GetOrder(ctx, id)
}
