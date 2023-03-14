package postgresq

import (
	"context"
	"fmt"

	"github.com/AxulReich/kitchen/internal/pkg/database"
	"github.com/AxulReich/kitchen/internal/pkg/domain"
	"github.com/AxulReich/kitchen/internal/repository"
)

type KitchenOrderRepo struct {
	db database.Ops
}

func NewKitchenOrderRepo(db database.Ops) *KitchenOrderRepo {
	return &KitchenOrderRepo{db: db}
}

func (r *KitchenOrderRepo) Create(ctx context.Context, order repository.KitchenOrder) (int64, error) {
	var id int64
	err := r.db.QueryRow(ctx, `
		INSERT INTO kitchen_order (
			shop_order_id,
            status
        ) VALUES ($1,$2)
		RETURNING id`,
		order.ShopOrderID,
		string(domain.KitchenOrderStatusNew),
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("can't execute query: %w", err)
	}

	return id, nil
}

func (r *KitchenOrderRepo) UpdateStatus(ctx context.Context, orderID int64, status string) error {
	result, err := r.db.Exec(ctx, `
		UPDATE kitchen_order
		SET
			status=$1
		WHERE id = $2`,
		status,
		orderID,
	)

	if err != nil {
		return fmt.Errorf("can't execute query: %w", err)
	}
	if result.RowsAffected() == 0 {
		return repository.ErrNoRowsAffected
	}

	return nil
}
