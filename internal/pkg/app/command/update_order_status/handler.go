package update_order_status

import (
	"context"
	"errors"
	"fmt"
	"github.com/AxulReich/kitchen/internal/pkg/domain"
	"github.com/Shopify/sarama"
	jsoniter "github.com/json-iterator/go"

	"github.com/AxulReich/kitchen/internal/pkg/kafka/kitchen_order_events_sender"
	"github.com/AxulReich/kitchen/internal/repository"
)

type UpdateOrderStatusHandler interface {
	Handle(ctx context.Context, command Command) error
}

type Handler struct {
	kitchenRepo repository.KitchenOrderRepository
	sender      kitchen_order_events_sender.Producer
}

func NewHandler(kitchenRepo repository.KitchenOrderRepository, sender kitchen_order_events_sender.Producer) *Handler {
	return &Handler{kitchenRepo: kitchenRepo, sender: sender}
}

func (h Handler) Handle(ctx context.Context, command Command) error {
	err := h.kitchenRepo.UpdateStatus(ctx, command.ShopOrderID, convertKitchenStatus(command.Status))
	if errors.Is(err, repository.ErrNoRowsAffected) {
		return nil
	}

	if err != nil {
		return err
	}

	go func() {
		data, _ := jsoniter.Marshal(domain.KitchenOrder{
			ShopOrderID: command.ShopOrderID,
			Status:      command.Status,
		})

		_ = h.sender.SendKitchenOrderStatusEvent(sarama.StringEncoder(fmt.Sprintf("%s", command.ShopOrderID)), data)
	}()

	return nil
}

func convertKitchenStatus(status domain.KitchenOrderStatusEnum) string {
	switch status {
	case domain.Values.CookingStart:
		return domain.Values.CookingStart.String()
	case domain.Values.New:
		return domain.Values.New.String()
	case domain.Values.CookingEnd:
		return domain.Values.CookingEnd.String()
	}
	return domain.Values.Invalid.String()
}
