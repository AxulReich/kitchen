package app

import (
	"github.com/AxulReich/kitchen/internal/pkg/app/command/update_order_status"
	"github.com/AxulReich/kitchen/internal/pkg/app/query/get_orders"
	"github.com/AxulReich/kitchen/internal/pkg/kafka/handler/shop_order"
	"github.com/AxulReich/kitchen/internal/pkg/kitchen_order"
	"github.com/AxulReich/kitchen/internal/repository/postgresq"
)

type handlerCollection struct {
	kafkaShopOrderHandler    *shop_order.Handler
	kitchenOrderHandler      *kitchen_order.Handler
	updateOrderStatusHandler *update_order_status.Handler
	getOrdersHandler         *get_orders.Handler
}

func (a *Application) initHandlers() {
	// TODO: clean the mess
	collection := handlerCollection{}

	collection.kitchenOrderHandler = kitchen_order.NewHandler(a.db, postgresq.Factory{})
	collection.kafkaShopOrderHandler = shop_order.NewHandler(collection.kitchenOrderHandler.Handle)
	collection.updateOrderStatusHandler = update_order_status.NewHandler(a.db, postgresq.Factory{}, a.messageSender)
	collection.getOrdersHandler = get_orders.NewHandler(a.db)

	a.handlers = collection
}
