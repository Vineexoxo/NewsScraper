package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/shishir54234/NewsScraper/backend/pkg/database"
	grpc1 "github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/configurations"
	"github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/features/scrape_service/v1/commands"
	"go.uber.org/fx"
	"go.opentelemetry.io/otel/trace"

	scraperpb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
)

func RunServers(
	lc fx.Lifecycle,
	log logger.ILogger,
	ctx context.Context,
	cfg *configurations.Config,
	rabbitmqPublisher rabbitmq.IPublisher,
	grpcServer *grpc1.GrpcServer,
	redisClient *database.RedisDB,
	tracer trace.Tracer,
) {
	// safety check
	if log == nil {
		panic("logger is nil")
	}

	lc.Append(fx.Hook{
		OnStart: func(startCtx context.Context) error {
			scraperHandler := commands.NewScraperServer(
				redisClient,
				rabbitmqPublisher,
				log,
				tracer,
				cfg.JobTTLSeconds,
			)

			scraperpb.RegisterScraperServiceServer(
				grpcServer.Grpc,
				scraperHandler,
			)

			go func(l logger.ILogger) {
				err := grpcServer.RunGrpcServer(ctx)
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					l.Panicf("error running grpc server: %v", err)
				}
			}(log) // pass logger explicitly to goroutine

			log.Infof("web-scraper gRPC server started on %s", cfg.Grpc.Port)
			return nil
		},
		OnStop: func(stopCtx context.Context) error {
			log.Infof("all servers shutdown gracefully...")
			return nil
		},
	})
}
