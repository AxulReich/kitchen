package update_order_status

import "github.com/AxulReich/kitchen/internal/pkg/domain"

type (
	Command struct {
		ShopOrderID int64
		Status      domain.KitchenOrderStatusEnum
	}

	ShopOrder struct {
		ID     int64  `json:"id"`
		Status string `json:"status"`
	}
)
