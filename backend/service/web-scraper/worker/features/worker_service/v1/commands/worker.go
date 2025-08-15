package commands

import (
	"context"
	"time"

	"github.com/gocolly/colly/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	pb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
	"github.com/streadway/amqp"
)

// WorkerDependencies holds shared services for the worker
type WorkerDependencies struct {
	Redis         *database.RedisDB
	Log           logger.ILogger
	ResultTTLSecs int // TTL in seconds
}

// NewWorkerDependencies is like a constructor for WorkerDependencies
func NewWorkerDependencies(redis *database.RedisDB, log logger.ILogger, ttlSecs int) *WorkerDependencies {
	return &WorkerDependencies{
		Redis:         redis,
		Log:           log,
		ResultTTLSecs: ttlSecs,
	}
}

// ScrapeJobHandler handles a single RabbitMQ message
func (w *WorkerDependencies) ScrapeJobHandler(queue string, delivery amqp.Delivery) error {
	var msg models.ScrapeJobMessage
	if err := jsoniter.Unmarshal(delivery.Body, &msg); err != nil {
		w.Log.Errorf("Failed to unmarshal message: %v", err)
		return err
	}

	jobID := msg.JobID
	ctx := context.Background()
	ttl := time.Duration(w.ResultTTLSecs) * time.Second

	// Initialize result with RUNNING status
	result := &pb.GetResultResponse{
		Status: pb.Status_RUNNING,
		Page: &pb.PageData{
			Url: msg.URL,
		},
	}

	if err := w.Redis.SetJobResult(ctx, jobID, msg.URL, result, ttl); err != nil {
		w.Log.Errorf("Failed to set initial job result: %v", err)
		return err
	}

	// Create Colly collector
	c := colly.NewCollector(colly.UserAgent(msg.UserAgent), colly.MaxDepth(1))
	c.OnHTML("title", func(e *colly.HTMLElement) { result.Page.Title = e.Text })
	c.OnHTML("body", func(e *colly.HTMLElement) { result.Page.Text = e.Text })
	c.OnError(func(r *colly.Response, err error) {
		w.Log.Errorf("Failed to scrape URL %s: %v", msg.URL, err)
		result.Status = pb.Status_FAILED
		_ = w.Redis.SetJobResult(ctx, jobID, msg.URL, result, ttl)
	})

	// Visit the URL
	if err := c.Visit(msg.URL); err != nil {
		w.Log.Errorf("Failed to visit URL %s: %v", msg.URL, err)
		result.Status = pb.Status_FAILED
		_ = w.Redis.SetJobResult(ctx, jobID, msg.URL, result, ttl)
		return err
	}

	// Mark job as COMPLETED
	result.Status = pb.Status_COMPLETED
	if err := w.Redis.SetJobResult(ctx, jobID, msg.URL, result, ttl); err != nil {
		w.Log.Errorf("Failed to set job result: %v", err)
		return err
	}

	w.Log.Infof("Successfully completed scraping job: %s", jobID)
	return nil
}

// StartWorker sets up the RabbitMQ consumer and starts processing
func StartWorker(ctx context.Context, rabbitCfg *rabbitmq.RabbitMQConfig, amqpConn *amqp.Connection, w *WorkerDependencies) error {
	handler := func(queue string, msg amqp.Delivery, deps interface{}) error {
		// Call the method on WorkerDependencies
		return w.ScrapeJobHandler(queue, msg)
	}

	consumer := rabbitmq.NewConsumer(ctx, rabbitCfg, amqpConn, w.Log, nil, handler)
	return consumer.ConsumeMessage(models.ScrapeJobMessage{}, w)
}
