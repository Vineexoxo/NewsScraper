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
	if req.GetJobId() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing job_id")
	}

	statusStr, err := s.redis.GetJobStatus(ctx, req.GetJobId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "job not found")
	}

	resp := &pb.GetResultResponse{JobId: req.GetJobId()}

	switch statusStr {
	case "QUEUED":
		resp.Status = pb.Status_QUEUED
	case "RUNNING":
		resp.Status = pb.Status_RUNNING
	case "COMPLETED":
		resp.Status = pb.Status_COMPLETED
		var page pb.PageData
		if err := s.redis.GetJobResult(ctx, req.GetJobId(), &page); err == nil {
			resp.Page = &page
		}
	case "FAILED":
		resp.Status = pb.Status_FAILED
	default:
		resp.Status = pb.Status_FAILED
		resp.Error = "unknown status"
	}

	// (optional) annotate trace with last snapshot
	if span := trace.SpanFromContext(ctx); span != nil {
		if b, err := jsoniter.Marshal(resp); err == nil {
			_ = b // attach if you want
		}
	}
	return resp, nil
}
