package postgresq

import (
	"github.com/AxulReich/kitchen/internal/pkg/database"
	"github.com/AxulReich/kitchen/internal/repository"
)

type Factory struct{}

func (f Factory) NewKitchenOrderRepository(ops database.Ops) repository.KitchenOrderRepository {
	return NewKitchenOrderRepo(ops)
}

func (f Factory) NewItemRepository(ops database.Ops) repository.ItemRepository {
	return NewItemRepo(ops)
}
