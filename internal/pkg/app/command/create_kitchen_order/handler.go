package create_kitchen_order

import (
	"context"
	"fmt"

	"github.com/AxulReich/kitchen/internal/pkg/domain"
	"github.com/AxulReich/kitchen/internal/repository"
)

type CreateKitchenOrderHandler interface {
	Handle(ctx context.Context, command Command) error
}

type Handler struct {
	kitchenRepo repository.KitchenOrderRepository
}

func NewHandler(kitchenRepo repository.KitchenOrderRepository) *Handler {
	return &Handler{kitchenRepo: kitchenRepo}
}

func (h *Handler) Handle(ctx context.Context, command Command) error {
	//	Create(ctx context.Context, order KitchenOrder, items ...Item) error
	var (
		orderDB = repository.KitchenOrder{
			ShopOrderID: command.Order.ShopOrderID,
			Status:      "new",
		}
		itemsDB = make([]repository.Item, 0, len(command.Order.Items))
	)
	for _, item := range command.Order.Items {
		itemsDB = append(itemsDB, repository.Item{
			Name:    item.Name,
			Comment: item.Comment,
		})
	}

	err := h.kitchenRepo.Create(ctx, orderDB, itemsDB)

	if err != nil {
		return fmt.Errorf("create_kitchen_order.Handler.Handle: %w", err)
	}

	return nil
}

func makeRepositoryItems(kitchenOrderID int64, items ...domain.Item) []repository.Item {
	result := make([]repository.Item, 0, len(items))

	for _, item := range items {
		result = append(result, repository.Item{
			KitchenOrderID: kitchenOrderID,
			Name:           item.Name,
			Comment:        item.Comment,
		})
	}

	return result
}
