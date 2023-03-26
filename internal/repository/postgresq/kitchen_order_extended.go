package postgresq

import (
	"context"

	"github.com/AxulReich/kitchen/internal/pkg/database"
	"github.com/AxulReich/kitchen/internal/repository"
	"github.com/georgysavva/scany/pgxscan"
)

type KitchenOrderExtendedRepo struct {
	db database.Ops
}

func NewKitchenOrderExtendedRepo(db database.Ops) *KitchenOrderExtendedRepo {
	return &KitchenOrderExtendedRepo{db: db}
}

func (r *KitchenOrderExtendedRepo) List(ctx context.Context, offset, limit int64) ([]repository.KitchenOrderExtended, int64, error) {
	var result []repository.KitchenOrderExtended

	err := pgxscan.Select(
		ctx,
		r.db,
		&result,
		`
			select orders.shop_order_id as id,
				   orders.status        as status,
				   item.item_name       as item_name,
				   item.item_comment    as item_comment
			from (select * from kitchen_order offset $1 limit $2) orders
			join item on orders.id = item.kitchen_order_id`,
		offset,
		limit,
	)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	err = r.db.QueryRow(ctx, `select count(*) as total from kitchen_order`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}
