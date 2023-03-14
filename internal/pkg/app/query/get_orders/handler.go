package get_orders

import (
	"context"

	"github.com/AxulReich/kitchen/internal/pkg/database"
)

type GetOrdersHandler interface {
	Handle(ctx context.Context, command Command) (Result, error)
}

type Handler struct {
	db database.DB
}

func NewHandler(db database.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Handle(ctx context.Context, command Command) (Result, error) {
	//	TODO: easy select with join
	return Result{}, nil
}
