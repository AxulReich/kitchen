package update_order_status

import (
	"context"

	"github.com/AxulReich/kitchen/internal/pkg/database"
	"github.com/AxulReich/kitchen/internal/pkg/kafka/sender"
	"github.com/AxulReich/kitchen/internal/repository"
)

type UpdateOrderStatusHandler interface {
	Handle(ctx context.Context, command Command) error
}

type Handler struct {
	db      database.DB
	factory repository.Factory

	sender sender.Producer //nolint:unused
}

func NewHandler(db database.DB, factory repository.Factory, sender sender.Producer) *Handler {
	return &Handler{db: db, factory: factory}
}

func (h Handler) Handle(ctx context.Context, command Command) error {

	return nil
}
