package app

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/AxulReich/kitchen/internal/config"
	"github.com/AxulReich/kitchen/internal/pkg/logger"
	"github.com/Shopify/sarama"
)

type kafkaWorker struct {
	shopOrderConsumer  sarama.ConsumerGroup
	handler            sarama.ConsumerGroupHandler
	topic              []string
	sessionRestartWait time.Duration
}

func newKafkaWorker(ctx context.Context, cfg *config.Config, handler sarama.ConsumerGroupHandler) (*kafkaWorker, error) {
	shopOrderListener, err := newKafkaConsumerGroup(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &kafkaWorker{
		shopOrderConsumer:  shopOrderListener,
		handler:            handler,
		topic:              []string{cfg.ShopOrderEventTopic},
		sessionRestartWait: cfg.SessionRestartWait,
	}, nil
}

func (w *kafkaWorker) work(ctx context.Context) {
	go func() {
		for {
			consumerCtx := context.Background()
			if err := w.shopOrderConsumer.Consume(consumerCtx, w.topic, w.handler); err != nil {
				logger.Error(ctx, fmt.Errorf("initalize shop order CG new session due to: %w", err))
			}
			time.Sleep(w.sessionRestartWait)
		}
	}()
}

func (w *kafkaWorker) close() error {
	return w.shopOrderConsumer.Close()
}

func newKafkaConsumerGroup(ctx context.Context, cfg *config.Config) (sarama.ConsumerGroup, error) {
	var (
		sessionTimeout    = cfg.SessionTimeout
		heartbeatInterval = cfg.HeartbeatInterval

		cfgVersion      = cfg.KafkaVersion
		balanceStrategy = sarama.BalanceStrategyRoundRobin
	)

	if (heartbeatInterval >= sessionTimeout) || (sessionTimeout/heartbeatInterval) < 3 {
		heartbeatInterval = sessionTimeout / 3
		logger.Warn(ctx, "heartbeatInterval was reset due to incorrect config")
	}

	consumerCfg := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion(cfgVersion)
	if err != nil {
		return nil, err
	}

	consumerCfg.Version = version
	consumerCfg.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumerCfg.Consumer.Group.Rebalance.Strategy = balanceStrategy
	consumerCfg.Consumer.Group.Rebalance.Retry.Backoff = cfg.RetryBackoff
	consumerCfg.Consumer.Group.Session.Timeout = cfg.SessionTimeout
	consumerCfg.Consumer.Group.Heartbeat.Interval = heartbeatInterval

	return sarama.NewConsumerGroup(getStrings(cfg.KafkaBrokers), cfg.GroupName, consumerCfg)
}

func newKafkaProducer(ctx context.Context, cfg *config.Config) (sarama.SyncProducer, error) {
	producerCfg := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion(cfg.KafkaVersion)
	if err != nil {
		return nil, err
	}
	producerCfg.Version = version
	// configure this
	producerCfg.Producer.Retry.Max = 5
	producerCfg.Producer.Retry.Backoff = 1 * time.Second
	// TODO: add metrics for success/error message sending and handle err chan
	producerCfg.Producer.Return.Successes = true
	producerCfg.Producer.Return.Errors = true

	return sarama.NewSyncProducer(getStrings(cfg.KafkaBrokers), producerCfg)
}

func getStrings(v string) []string {
	if len(v) == 0 {
		return nil
	}
	return strings.Split(v, ",")
}
