package app

import "github.com/AxulReich/kitchen/internal/repository/postgresq"

func (a *Application) initRepositoryCollection() {
	a.repositories = repositoryCollection{
		kitchenOrderRepository:         postgresq.NewKitchenOrderRepo(a.db),
		kitchenOrderExtendedRepository: postgresq.NewKitchenOrderExtendedRepo(a.db),
	}
}
