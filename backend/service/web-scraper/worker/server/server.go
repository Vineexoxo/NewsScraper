package worker

import (
	"context"

	"github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/web-scraper/worker/features/worker_service/v1/commands"
	"github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/configurations"
	"github.com/streadway/amqp"
	"go.uber.org/fx"
)

// RunWorker integrates the worker into the FX lifecycle
func RunWorker(
	lc fx.Lifecycle,
	log logger.ILogger,
	ctx context.Context,
	rabbitCfg *rabbitmq.RabbitMQConfig,
	amqpConn *amqp.Connection,
	redis *database.RedisDB,
	cfg *configurations.Config,
) {
	// Use constructor to create WorkerDependencies
	deps := commands.NewWorkerDependencies(redis, log, cfg.JobTTLSeconds)

	lc.Append(fx.Hook{
		OnStart: func(startCtx context.Context) error {
			log.Infof("Starting web-scraper worker...")
			go func() {
				if err := commands.StartWorker(ctx, rabbitCfg, amqpConn, deps); err != nil {
					log.Errorf("Worker failed: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(stopCtx context.Context) error {
			log.Infof("Worker shutdown gracefully...")
			return nil
		},
	})
}
