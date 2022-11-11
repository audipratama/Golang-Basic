package _interface

import (
	"context"
	"golang-basic/model"
)

type OrderRepository interface {
	Insert(ctx context.Context, order model.Orders) (int64, error)
	GetOrdersByIDs(ctx context.Context, id []int64) ([]model.Orders, error)
	UpdateOrderByID(ctx context.Context, order model.Orders) (model.Orders, error)
}
