package commands

import (
	"context"

	jsoniter "github.com/json-iterator/go"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
)

// GetResult retrieves the status and result of a scraping job by job ID
// It checks the job status in Redis and returns the appropriate response
func (s *ScraperServer) GetResult(ctx context.Context, req *pb.GetResultRequest) (*pb.GetResultResponse, error) {
	s.log.Infof("Received GetResult request: JobID=%s", req.GetJobId())

	if req.GetJobId() == "" {
		s.log.Infof("GetResult request missing job_id")
		return nil, status.Error(codes.InvalidArgument, "missing job_id")
	}

	result, err := s.redis.GetJobResult(ctx, req.GetJobId())
	if err != nil {
		s.log.Errorf("Failed to get job result from Redis for JobID=%s: %v", req.GetJobId(), err)
		return nil, status.Errorf(codes.NotFound, "job not found")
	}

	s.log.Infof("Successfully retrieved result for JobID=%s, Status=%s", req.GetJobId(), result.GetStatus().String())

	// Optional: annotate trace
	if span := trace.SpanFromContext(ctx); span != nil {
		if b, err := jsoniter.Marshal(result); err == nil {
			_ = b // could attach to span if desired
		} else {
			s.log.Infof("Failed to marshal result for tracing: %v", err)
		}
	}

	return result, nil
}
