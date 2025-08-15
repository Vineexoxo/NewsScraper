package grpcclient

import (
	"context"
	"fmt"
	"time"

	pb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
	"google.golang.org/grpc"
)

type WebScraperClient struct {
	client pb.ScraperServiceClient
}
type PageData struct {
	Url string 
	Title string
	Text  string
}

func NewWebScraperClient(connAddr string) *WebScraperClient {
	conn, err:= grpc.Dial(connAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to dial server:", err)
		return nil	
	}	

	client := pb.NewScraperServiceClient(conn)
	return &WebScraperClient{
		client: client,
	}
}


// give it the string it will interact with the client and get the description 
func (c* WebScraperClient) CallTheClient(ctx context.Context, url string) (*PageData, error){
	curr:= &pb.ScrapeRequest{Url: url}	

	resp, err:= c.client.ScrapePage(ctx, curr)
	if err!=nil{
		return nil, err
	}
	res, err := c.client.GetResult(ctx, &pb.GetResultRequest{JobId: resp.GetJobId()})
	if err != nil {
		return nil, fmt.Errorf("initial GetResult failed: %w", err)
	}

	for res.Status != pb.Status_COMPLETED {
		// Optional: check for failure state
		if res.Status == pb.Status_FAILED {
			return nil, fmt.Errorf("job %s failed", resp.GetJobId())
		}

		// Wait before polling again
		time.Sleep(2 * time.Second)

		// Respect context cancel/deadline
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// Call GetResult again
		res, err = c.client.GetResult(ctx, &pb.GetResultRequest{JobId: resp.GetJobId()})
		if err != nil {
			return nil, fmt.Errorf("polling GetResult failed: %w", err)
		}
	}

	return &PageData{Url: res.GetPage().GetUrl(), Title: res.GetPage().Title, Text: res.GetPage().Text}, nil
}

func (c *WebScraperClient) ScrapePage(ctx context.Context, req *pb.ScrapeRequest) (*pb.ScrapeResponse, error) {
	return c.client.ScrapePage(ctx, req)
}

func (c *WebScraperClient) GetResult(ctx context.Context, req *pb.GetResultRequest) (*pb.GetResultResponse, error) {
	return c.client.GetResult(ctx, req)
}











