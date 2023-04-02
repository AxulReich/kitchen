package update_order_status

import (
	"context"
	"fmt"
	"github.com/AxulReich/kitchen/internal/pkg/domain"
	"github.com/Shopify/sarama"

	"github.com/AxulReich/kitchen/internal/pkg/kafka/sender"
	"github.com/AxulReich/kitchen/internal/repository"
)

type UpdateOrderStatusHandler interface {
	Handle(ctx context.Context, command Command) error
}

type Handler struct {
	kitchenRepo repository.KitchenOrderRepository
	sender      sender.Producer //nolint:unused
}

func NewHandler(kitchenRepo repository.KitchenOrderRepository, sender sender.Producer) *Handler {
	return &Handler{kitchenRepo: kitchenRepo, sender: sender}
}

func (h Handler) Handle(ctx context.Context, command Command) error {
	err := h.kitchenRepo.UpdateStatus(ctx, command.KitchenOrderID, convertKitchenStatus(command.Status))
	if err != nil {
		return err
	}

	go func() {
		_ = h.sender.SendMessage(sarama.StringEncoder(fmt.Sprintf("%s", command.KitchenOrderID)), []byte(convertKitchenStatus(command.Status)))
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
