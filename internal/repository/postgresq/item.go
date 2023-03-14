package postgresq

import (
	"context"
	"fmt"

	"github.com/AxulReich/kitchen/internal/pkg/database"
	"github.com/AxulReich/kitchen/internal/repository"
)

type ItemRepo struct {
	db database.Ops
}

func NewItemRepo(db database.Ops) *ItemRepo {
	return &ItemRepo{db: db}
}

func (r *ItemRepo) Create(ctx context.Context, items ...repository.Item) error {
	if len(items) == 0 {
		return nil
	}

	query, args := makeLabelGroupCreateQuery(items)
	_, err := r.db.Query(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("can't execute query: %w", err)
	}

	return nil
}

func makeLabelGroupCreateQuery(items []repository.Item) (string, []interface{}) {
	valueArgs := make([]interface{}, 0, len(items)*4)

	for _, item := range items {
		valueArgs = append(valueArgs, item.KitchenOrderID)
		valueArgs = append(valueArgs, item.Name)
		valueArgs = append(valueArgs, item.Comment)
	}

	sqlStr := database.GetBulkInsertSQL("item", []string{"kitchen_order_id", "item_name", "item_comment"}, len(items))
	return sqlStr, valueArgs
}
