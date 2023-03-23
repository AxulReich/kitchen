package get_orders

import (
	"context"
	"github.com/AxulReich/kitchen/internal/pkg/domain"
	"github.com/AxulReich/kitchen/internal/repository"
)

type GetOrdersHandler interface {
	Handle(ctx context.Context, command Command) (Result, error)
}

type Handler struct {
	kitchenOrderExtended repository.KitchenOrderExtendedRepository
}

func NewHandler(kitchenOrderExtended repository.KitchenOrderExtendedRepository) *Handler {
	return &Handler{kitchenOrderExtended: kitchenOrderExtended}
}

func (h *Handler) Handle(ctx context.Context, command Command) (Result, error) {
	res, total, err := h.kitchenOrderExtended.List(ctx, command.Offset, command.Limit)
	if err != nil {
		return Result{}, err
	}

	orders := make(map[int64][]domain.Item)
	ordersStatus := make(map[int64]string)
	for _, order := range res {
		orders[order.ShopOrderID] = append(orders[order.ShopOrderID], domain.Item{
			Name:    order.ItemName,
			Comment: order.ItemComment,
		})
		ordersStatus[order.ShopOrderID] = order.Status
	}

	var ordersRes []domain.KitchenOrder
	for id, items := range orders {
		ordersRes = append(ordersRes, domain.KitchenOrder{
			ShopOrderID: id,
			//Status:      ordersStatus[id],
			Items: items,
		})
	}

	return Result{
		Orders: ordersRes,
		Total:  total,
	}, nil
}
