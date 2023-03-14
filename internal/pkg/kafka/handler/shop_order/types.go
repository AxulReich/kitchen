package shop_order

type (
	Item struct {
		Name    string `json:"name"`
		Comment string `json:"comment"`
	}

	ShopOrderMessage struct {
		ID     int64  `json:"id"`
		Status string `json:"status"`
		Items  []Item `json:"items"`
	}
)
