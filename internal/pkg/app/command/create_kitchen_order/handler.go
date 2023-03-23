package create_kitchen_order

import (
	"context"
	"fmt"

	"github.com/AxulReich/kitchen/internal/pkg/database"
	"github.com/AxulReich/kitchen/internal/pkg/domain"
	"github.com/AxulReich/kitchen/internal/repository"
	"github.com/jackc/pgx/v4"
)

type CreateKitchenOrderHandler interface {
	Handle(ctx context.Context, command Command) error
}

type Handler struct {
	db      database.DB
	factory repository.Factory
}

func NewHandler(db database.DB, factory repository.Factory) *Handler {
	return &Handler{db: db, factory: factory}
}

func (h *Handler) Handle(ctx context.Context, command Command) error {
	err := h.db.WithTx(ctx, func(tx database.Ops) error {
		var (
			orderRepo = h.factory.NewKitchenOrderRepository(tx)
			itemRepo  = h.factory.NewItemRepository(tx)
		)

		kitchenOrderID, err := orderRepo.Create(ctx, repository.KitchenOrder{
			ShopOrderID: command.Order.ShopOrderID,
			Status:      fmt.Sprintf("%s", command.Order.Status),
		})

		if err != nil {
			return fmt.Errorf("orderRepository.Create: %w", err)
		}

		err = itemRepo.Create(ctx, makeRepositoryItems(kitchenOrderID, command.Order.Items...)...)
		if err != nil {
			return fmt.Errorf("itemRepository.Create: %w", err)
		}

		return nil

	}, database.ReadWrite(), database.IsolationLevel(pgx.RepeatableRead))

	if err != nil {
		return fmt.Errorf("tx: %w", err)
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
