package shop_order

import (
	"context"
	"fmt"

	"github.com/AxulReich/kitchen/internal/pkg/app/command/create_kitchen_order"

	"github.com/AxulReich/kitchen/internal/pkg/domain"
	"github.com/AxulReich/kitchen/internal/pkg/logger"
	"github.com/Shopify/sarama"
	jsoniter "github.com/json-iterator/go"
)

type Handler struct {
	handler create_kitchen_order.CreateKitchenOrderHandler
}

func NewHandler(handler create_kitchen_order.CreateKitchenOrderHandler) *Handler {
	return &Handler{handler: handler}
}

// TODO: add metrics and alert for invalid messages
func (h *Handler) process(ctx context.Context, msg *sarama.ConsumerMessage) error {
	var shopOrderMsg ShopOrderMessage
	if err := jsoniter.Unmarshal(msg.Value, &shopOrderMsg); err != nil {
		logger.Errorf(ctx, "can't unmurshal message: %s err: %w", string(msg.Value), err)
		return nil
	}

	if shopOrderMsg.ID == 0 {
		return fmt.Errorf("invalid shop order status: %d", shopOrderMsg.ID)
	}
	if len(shopOrderMsg.Items) == 0 {
		return fmt.Errorf("invalid shop order items len: %d", shopOrderMsg.ID)
	}

	if shopOrderMsg.Status != string(domain.ShopOrderStatusConfirmed) {
		return fmt.Errorf("invalid shop order items len: %s", shopOrderMsg.Status)
	}

	kitchenOrder := domain.KitchenOrder{
		ShopOrderID: shopOrderMsg.ID,
		Status:      domain.Values.New,
	}

	items := make([]domain.Item, 0, len(shopOrderMsg.Items))
	for _, item := range shopOrderMsg.Items {
		items = append(items, domain.Item{
			Name:    item.Name,
			Comment: item.Comment,
		})
	}
	kitchenOrder.Items = items

	if err := h.handler.Handle(ctx, create_kitchen_order.Command{Order: kitchenOrder}); err != nil {
		return err
	}

	return nil
}

func (h *Handler) Cleanup(cgSession sarama.ConsumerGroupSession) error {
	return nil
}

func (h *Handler) Setup(cgSession sarama.ConsumerGroupSession) error {
	logger.Errorf(context.Background(), "rebalancing for partitions: %v", cgSession.Claims())
	return nil
}

// ConsumeClaim calls handler for every kafka message that arrived from topic
func (h *Handler) ConsumeClaim(session sarama.ConsumerGroupSession, cgClaim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-cgClaim.Messages():
			ctx := context.Background()
			if !ok {
				logger.Infof(ctx, "shop-order-event read channel closed")
				return nil
			}

			err := h.process(ctx, msg)

			if err != nil {
				// this return will stop consuming from claim and close consumer session
				err = fmt.Errorf("shop-order-event ConsumeClaim error after handle: %w", err)
				logger.Errorf(ctx, "%s; Message that provoke error: %s", err.Error(), string(msg.Value))
				return err
			}

			// NOTE: if error occurs then we will start to try to consume message that provoked an error
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			logger.Infof(context.Background(), "shop-order-event session context done")
			return nil
		}
	}
}
