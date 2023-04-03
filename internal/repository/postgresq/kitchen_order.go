package postgresq

import (
	"context"
	"fmt"

	"github.com/AxulReich/kitchen/internal/pkg/database"
	"github.com/AxulReich/kitchen/internal/repository"
	"github.com/jackc/pgx/v4"
)

type KitchenOrderRepo struct {
	db database.DB
}

func NewKitchenOrderRepo(db database.DB) *KitchenOrderRepo {
	return &KitchenOrderRepo{db: db}
}

func (r *KitchenOrderRepo) Create(ctx context.Context, order repository.KitchenOrder, items []repository.Item) error {
	err := r.db.WithTx(ctx, func(tx database.Ops) error {
		kitchenOrderID, err := r.createOrder(ctx, order.ShopOrderID, order.Status)

		if err != nil {
			return fmt.Errorf("orderRepository.Create: %w", err)
		}

		err = r.createItems(ctx, kitchenOrderID, items)
		if err != nil {
			return fmt.Errorf("itemRepository.Create: %w", err)
		}

		return nil

	}, database.ReadWrite(), database.IsolationLevel(pgx.ReadCommitted))

	if err != nil {
		return fmt.Errorf("KitchenOrderRepo.Create: %w", err)
	}

	return nil
}

func (r *KitchenOrderRepo) createOrder(ctx context.Context, orderID int64, status string) (int64, error) {
	var id int64

	err := r.db.QueryRow(ctx, `
		INSERT INTO kitchen_order (
			shop_order_id,
            status
        ) VALUES ($1,$2)
		RETURNING id`,
		orderID,
		status,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("can't execute query: %w", err)
	}

	return id, nil
}

func (r *KitchenOrderRepo) createItems(ctx context.Context, kitchenOrderID int64, items []repository.Item) error {
	if len(items) == 0 {
		return nil
	}

	query, args := makeItemCreateCreateQuery(kitchenOrderID, items)
	_, err := r.db.Query(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("can't execute query: %w", err)
	}

	return nil
}

func makeItemCreateCreateQuery(kitchenOrderID int64, items []repository.Item) (string, []interface{}) {
	valueArgs := make([]interface{}, 0, len(items)*3)

	for _, item := range items {
		valueArgs = append(valueArgs, kitchenOrderID)
		valueArgs = append(valueArgs, item.Name)
		valueArgs = append(valueArgs, item.Comment)
	}

	sqlStr := database.GetBulkInsertSQL("item", []string{"kitchen_order_id", "item_name", "item_comment"}, len(items))
	return sqlStr, valueArgs
}

func (r *KitchenOrderRepo) UpdateStatus(ctx context.Context, orderID int64, status string) error {
	//TODO: first check if order exist
	result, err := r.db.Exec(ctx, `
		UPDATE kitchen_order
		SET
			status=$1
		WHERE shop_order_id = $2`,
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
