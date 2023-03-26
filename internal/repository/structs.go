package repository

type (
	KitchenOrderExtended struct {
		ShopOrderID int64  `db:"id"`
		Status      string `db:"status"`
		ItemName    string `db:"item_name"`
		ItemComment string `db:"item_comment"`
	}

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
