package grpcclient

import (
	pb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
	"context"
)

type WebScraperClient struct {
	pb.ScraperServiceClient
}


func NewWebScraperClient() *WebScraperClient {
	return &WebScraperClient{}
}


func (c *WebScraperClient) ScrapePage(ctx context.Context, req *pb.ScrapeRequest) (*pb.ScrapeResponse, error) {
	
	
	
	
	
	
	return c.ScrapePage(ctx, req)
}











