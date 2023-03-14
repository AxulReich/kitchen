package repository

import "github.com/AxulReich/kitchen/internal/pkg/database"

type Factory interface {
	NewKitchenOrderRepository(ops database.Ops) KitchenOrderRepository
	NewItemRepository(ops database.Ops) ItemRepository
}
