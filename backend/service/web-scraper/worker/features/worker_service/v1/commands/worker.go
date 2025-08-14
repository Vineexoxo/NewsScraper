package commands

import (
	"context"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/json-iterator/go"
	"github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	pb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

// WorkerDependencies holds shared services
type WorkerDependencies struct {
	Redis *database.RedisDB
	Log   logger.ILogger
}

// ScrapeJobHandler handles a single RabbitMQ message
func ScrapeJobHandler(queue string, delivery amqp.Delivery, deps WorkerDependencies) error {
	var msg models.ScrapeJobMessage
	if err := jsoniter.Unmarshal(delivery.Body, &msg); err != nil {
		deps.Log.Errorf("Failed to unmarshal message: %v", err)
		return err
	}

	jobID := msg.JobID
	ctx := context.Background()
	ttl := 1 * time.Hour

	if err := deps.Redis.SetJobStatus(ctx, jobID, "RUNNING", ttl); err != nil {
		deps.Log.Errorf("Failed to set job status RUNNING: %v", err)
		return err
	}

	c := colly.NewCollector(colly.UserAgent(msg.UserAgent), colly.MaxDepth(1))
	var pageData pb.PageData
	pageData.Url = msg.URL

	c.OnHTML("title", func(e *colly.HTMLElement) { pageData.Title = e.Text })
	c.OnHTML("body", func(e *colly.HTMLElement) { pageData.Text = e.Text })
	c.OnError(func(r *colly.Response, err error) {
		deps.Log.Errorf("Failed to scrape URL %s: %v", msg.URL, err)
		_ = deps.Redis.SetJobStatus(ctx, jobID, "FAILED", ttl)
	})

	if err := c.Visit(msg.URL); err != nil {
		deps.Log.Errorf("Failed to visit URL %s: %v", msg.URL, err)
		_ = deps.Redis.SetJobStatus(ctx, jobID, "FAILED", ttl)
		return err
	}

	if err := deps.Redis.SetJobResult(ctx, jobID, pageData, ttl); err != nil {
		deps.Log.Errorf("Failed to set job result: %v", err)
		return err
	}

	if err := deps.Redis.SetJobStatus(ctx, jobID, "COMPLETED", ttl); err != nil {
		deps.Log.Errorf("Failed to set job status COMPLETED: %v", err)
		return err
	}

	deps.Log.Infof("Successfully completed scraping job: %s", jobID)
	return nil
}

// StartWorker sets up the RabbitMQ consumer and starts processing
func StartWorker(ctx context.Context, rabbitCfg *rabbitmq.RabbitMQConfig, amqpConn *amqp.Connection, deps WorkerDependencies) error {
	handler := func(queue string, msg amqp.Delivery, deps WorkerDependencies) error {
		return ScrapeJobHandler(queue, msg, deps)
	}

	consumer := rabbitmq.NewConsumer(ctx, rabbitCfg, amqpConn, deps.Log,nil, handler)
	return consumer.ConsumeMessage(models.ScrapeJobMessage{}, deps)
}
