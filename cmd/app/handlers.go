package app

import (
	"github.com/AxulReich/kitchen/internal/pkg/app/command/create_kitchen_order"
	"github.com/AxulReich/kitchen/internal/pkg/app/command/update_order_status"
	"github.com/AxulReich/kitchen/internal/pkg/app/query/get_orders"
	"github.com/AxulReich/kitchen/internal/pkg/kafka/handler/shop_order"
)

type handlerCollection struct {
	kafkaShopOrderHandler    *shop_order.Handler
	kitchenOrderHandler      *create_kitchen_order.Handler
	updateOrderStatusHandler *update_order_status.Handler
	getOrdersHandler         *get_orders.Handler
}

func (a *Application) initHandlers() {
	// TODO: clean the mess
	collection := handlerCollection{}

	collection.kitchenOrderHandler = create_kitchen_order.NewHandler(a.repositories.kitchenOrderRepository)
	collection.kafkaShopOrderHandler = shop_order.NewHandler(collection.kitchenOrderHandler)

	collection.updateOrderStatusHandler = update_order_status.NewHandler(a.repositories.kitchenOrderRepository, a.messageSender)
	collection.getOrdersHandler = get_orders.NewHandler(a.repositories.kitchenOrderExtendedRepository)

	a.handlers = collection
}
