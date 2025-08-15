package commands

import (
	"context"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	pb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
	"github.com/streadway/amqp"

	"regexp"
	"strings"
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

// ExtractCleanTextFromHTML extracts readable text from HTML, removes scripts, ads, JSON, and returns chunks.
func ExtractCleanTextFromHTML(html string, windowSize, overlap int) []string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil
	}

	// Remove scripts, styles, noscript
	doc.Find("script, style, noscript").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	// Remove JSON-LD or embedded structured data
	doc.Find(`script[type="application/ld+json"]`).Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	// Remove known ad or widget containers
	adSelectors := []string{".taboola", ".adslot", ".vj_tearntake", "#taboola-right-rail-thumbnails-new-world", ".vuukle"}
	for _, sel := range adSelectors {
		doc.Find(sel).Each(func(i int, s *goquery.Selection) {
			s.Remove()
		})
	}

	var sb strings.Builder

	// Extract headings, paragraphs, list items
	doc.Find("h1,h2,h3,h4,h5,h6,p,li").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			// add markdown-style for list items
			if s.Is("li") {
				sb.WriteString("- " + text + "\n")
			} else {
				sb.WriteString(text + "\n\n")
			}
		}
	})

	// Normalize whitespace
	cleaned := normalizeWhitespace(sb.String())

	// Chunk into sliding window
	return chunkTextSlidingWindow(cleaned, windowSize, overlap)
}

// normalizeWhitespace replaces multiple whitespace/newlines with a single space
func normalizeWhitespace(text string) string {
	spaceRegex := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(spaceRegex.ReplaceAllString(text, " "))
}

// chunkTextSlidingWindow splits text into overlapping chunks
func chunkTextSlidingWindow(text string, windowSize, overlap int) []string {
	var chunks []string
	runes := []rune(text) // handle unicode
	start := 0

	for start < len(runes) {
		end := start + windowSize
		if end > len(runes) {
			end = len(runes)
		}

		chunk := strings.TrimSpace(string(runes[start:end]))
		if chunk != "" {
			chunks = append(chunks, chunk)
		}

		start += windowSize - overlap
		if start < 0 {
			start = 0
		}
	}

	return chunks
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
	c.OnHTML("head > title", func(e *colly.HTMLElement) {
		result.Page.Title = strings.TrimSpace(e.Text)
	})
	c.OnHTML("body", func(e *colly.HTMLElement) {
		html, _ := e.DOM.Html()
		chunks := ExtractCleanTextFromHTML(html, 1000, 0) // returns []string
		result.Page.Text = strings.Join(chunks, "\n\n")   // join chunks as one string
	})

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
