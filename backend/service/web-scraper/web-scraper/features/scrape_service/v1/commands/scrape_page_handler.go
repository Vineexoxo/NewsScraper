package commands

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	pb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
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

func (s *ScraperServer) ScrapePage(ctx context.Context, req *pb.ScrapeRequest) (*pb.ScrapeResponse, error) {
	if req.GetUrl() == ""  {
		return nil, status.Error(codes.InvalidArgument, "invalid url")
	}
	jobID := uuid.NewV4().String()

	ttl := time.Duration(s.cfgTTLSeconds) * time.Second
	if err := s.redis.SetJobStatus(ctx, jobID, "QUEUED", ttl); err != nil {
		return nil, status.Errorf(codes.Internal, "redis error: %v", err)
	}

	msg := &models.ScrapeJobMessage{
		JobID:     jobID,
		URL:       req.GetUrl(),
		UserAgent: req.GetUserAgent(),
	}
	if err := s.publisher.PublishMessage(msg); err != nil {
		_ = s.redis.SetJobStatus(ctx, jobID, "FAILED", ttl)
		return nil, status.Errorf(codes.Internal, "publish error: %v", err)
	}

	return &pb.ScrapeResponse{JobId: jobID, Status: pb.Status_QUEUED}, nil
}