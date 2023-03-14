package get_orders

import "github.com/AxulReich/kitchen/internal/pkg/domain"

type (
	Command struct {
		Offset int64
		Limit  int64
	}

	Result struct {
		Orders []domain.KitchenOrder
		Total  int64
	}
)
