package app

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/AxulReich/kitchen/internal/app/server"
	"github.com/AxulReich/kitchen/internal/config"
	"github.com/AxulReich/kitchen/internal/pkg/database"
	"github.com/AxulReich/kitchen/internal/pkg/kafka/kitchen_order_events_sender"
	"github.com/AxulReich/kitchen/internal/pkg/logger"
	"github.com/AxulReich/kitchen/internal/repository/postgresq"
)

type repositoryCollection struct {
	kitchenOrderRepository         *postgresq.KitchenOrderRepo
	kitchenOrderExtendedRepository *postgresq.KitchenOrderExtendedRepo
}

type closeError struct {
	errors []string
}

func (e closeError) Error() string {
	return strings.Join(e.errors, "; ")
}

type Application struct {
	db     database.DB
	worker *kafkaWorker

	repositories repositoryCollection
	handlers     handlerCollection

	messageSender *kitchen_order_events_sender.MessageSender
	server        *server.KitchenServer
	k8s           *http.Server
	shutDownChan  chan struct{}
}

func NewApplication(ctx context.Context, cfg *config.Config, shutDownChan chan struct{}) (*Application, error) {
	a := &Application{
		shutDownChan: shutDownChan,
	}

	{
		r := Router(ctx)
		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppHTTPK8SPort),
			Handler: r,
		}
		a.k8s = srv
	}

	{
		db, err := database.NewDB(ctx, database.Options{DSN: cfg.DbDsnEntrypoint})
		if err != nil {
			return nil, err
		}
		a.db = db
	}

	{
		producer, err := newKafkaProducer(ctx, cfg)
		if err != nil {
			return nil, err
		}
		a.messageSender = kitchen_order_events_sender.NewMessageSender(producer, cfg.KitchenOrderEventTopic)
	}

	a.initRepositoryCollection()
	a.initHandlers()

	{
		worker, err := newKafkaWorker(ctx, cfg, a.handlers.kafkaShopOrderHandler)
		if err != nil {
			return nil, err
		}

		a.worker = worker
	}

	a.server = server.NewServer(cfg, a.handlers.updateOrderStatusHandler, a.handlers.getOrdersHandler)

	return a, nil
}

func (a *Application) Run(ctx context.Context) error {
	go func() {
		err := a.k8s.ListenAndServe()
		if err != nil {
			a.shutDownChan <- struct{}{}
			logger.Error(ctx, "%v", err)
		}
	}()

	a.worker.work(ctx)

	return a.server.Run(ctx)
}

func (a *Application) Close() error {
	var closeErr = closeError{}

	if err := a.k8s.Shutdown(context.Background()); err != nil {
		closeErr.errors = append(closeErr.errors, err.Error())
	}

	if err := a.worker.close(); err != nil {
		closeErr.errors = append(closeErr.errors, err.Error())
	}

	if err := a.messageSender.Close(); err != nil {
		closeErr.errors = append(closeErr.errors, err.Error())
	}

	a.server.Stop()
	a.db.Close()

	return closeErr
}
