package commands

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	pb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
)

type ScraperServer struct {
	pb.UnimplementedScraperServiceServer

	cfgTTLSeconds int
	redis         *database.RedisDB
	publisher     rabbitmq.IPublisher
	log           logger.ILogger
	tracer        trace.Tracer
}

func NewScraperServer(
	redis *database.RedisDB,
	publisher rabbitmq.IPublisher,
	log logger.ILogger,
	tracer trace.Tracer,
	resultTTLSeconds int,
) *ScraperServer {
	return &ScraperServer{
		redis:         redis,
		publisher:     publisher,
		log:           log,
		tracer:        tracer,
		cfgTTLSeconds: resultTTLSeconds,
	}
}

// ScrapePage grpc ingests the request, publishes a message to RabbitMQ, and returns a job ID
// The job ID can be used to track the status of the scraping job
func (s *ScraperServer) ScrapePage(ctx context.Context, req *pb.ScrapeRequest) (*pb.ScrapeResponse, error) {
	if req.GetUrl() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid url")
	}

	ttl := time.Duration(s.cfgTTLSeconds) * time.Second
	// 1. Check if this URL is already cached
	_, statusVal, err := s.redis.GetResultByURL(ctx, req.GetUrl())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "redis get by URL error: %v", err)
	}
	if statusVal == pb.Status_COMPLETED {
		s.log.Infof("URL already cached, returning existing job ID")
		// Already in cache → return existing job ID
		existingJobID, err := s.redis.GetJobIDByURL(ctx, req.GetUrl())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "redis get jobID error: %v", err)
		}
		return &pb.ScrapeResponse{
			JobId:  existingJobID,
			Status: statusVal,
		}, nil
	}
	// 2. Create new job
	s.log.Infof("Creating new job for URL: %s", req.GetUrl())
	jobID := uuid.NewV4().String()
	result := &pb.GetResultResponse{
		JobId:  jobID,
		Status: pb.Status_QUEUED,
		Page:   nil,
	}

	// Store both job result and URL mapping in one go
	if err := s.redis.SetJobResult(ctx, jobID, req.GetUrl(), result, ttl); err != nil {
		return nil, status.Errorf(codes.Internal, "redis set job result error: %v", err)
	}

	// 3. Publish to RabbitMQ
	msg := &models.ScrapeJobMessage{
		JobID:     jobID,
		URL:       req.GetUrl(),
		UserAgent: req.GetUserAgent(),
	}
	if err := s.publisher.PublishMessage(msg); err != nil {
		// Mark job as failed in Redis
		result.Status = pb.Status_FAILED
		_ = s.redis.SetJobResult(ctx, jobID, req.GetUrl(), result, ttl)
		return nil, status.Errorf(codes.Internal, "publish error: %v", err)
	}

	return &pb.ScrapeResponse{JobId: jobID, Status: pb.Status_QUEUED}, nil
}
