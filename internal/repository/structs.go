package repository

type (
	KitchenOrder struct {
		ID          int64  `db:"id"`
		ShopOrderID int64  `db:"shop_order_id"`
		Status      string `db:"status"`
	}

	Item struct {
		KitchenOrderID int64  `db:"kitchen_order_id"`
		Name           string `db:"item_name"`
		Comment        string `db:"item_comment"`
	}
)
