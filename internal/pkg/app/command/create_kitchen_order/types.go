package create_kitchen_order

import "github.com/AxulReich/kitchen/internal/pkg/domain"

type Command struct {
	Order domain.KitchenOrder
}
