package app

import "github.com/AxulReich/kitchen/internal/repository/postgresq"

func (a *Application) initRepositoryCollection() {
	a.repositories = repositoryCollection{
		itemRepository:                 postgresq.NewItemRepo(a.db),
		kitchenOrderRepository:         postgresq.NewKitchenOrderRepo(a.db),
		kitchenOrderExtendedRepository: postgresq.NewKitchenOrderExtendedRepo(a.db),
	}
}
