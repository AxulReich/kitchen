package repository

import (
	"context"
	"errors"
)

// ErrNoRowsAffected must be returned from repository when modification request
// does not cause any changes
var ErrNoRowsAffected = errors.New("no rows affected")

type KitchenOrderRepository interface {
	Create(ctx context.Context, order KitchenOrder) (int64, error)
	UpdateStatus(ctx context.Context, orderID int64, status string) error
}

type ItemRepository interface {
	Create(ctx context.Context, items ...Item) error
}
