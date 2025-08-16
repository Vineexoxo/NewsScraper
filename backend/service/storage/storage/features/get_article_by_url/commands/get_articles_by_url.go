package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/contracts"
	dtos "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/get_article_by_url/dtos"
	grpcclient "github.com/shishir54234/NewsScraper/backend/service/storage/storage/grpc_client"
)

type GetArticlesByUrlHandler struct {
	log               logger.ILogger
	rabbitmqPublisher rabbitmq.IPublisher
	articleRepository contracts.ArticleRepository 
	ctx               context.Context
	web_scraper_client grpcclient.WebScraperClient
	llm_client grpcclient.LLMClient
}


func NewGetArticlesByUrlHandler (log logger.ILogger, rabbitmqPublisher *rabbitmq.IPublisher, 
articleRepository contracts.ArticleRepository, web_scraper_client grpcclient.WebScraperClient, 
llm_client grpcclient.LLMClient,
ctx context.Context) *GetArticlesByUrlHandler {
	return &GetArticlesByUrlHandler{log: log, rabbitmqPublisher: *rabbitmqPublisher, 
	web_scraper_client: web_scraper_client,
	llm_client: llm_client,
	articleRepository: articleRepository, ctx: ctx}
}



func (q *GetArticlesByUrlHandler) Handle(ctx context.Context, request dtos.RequestArticleDto) (
	*models.Article, error) {
	fmt.Println("MOTHERRFUCKER", request.URL)
	
	res,err:=q.articleRepository.GetArticleByUrl(q.ctx, request.URL)
	if err!=nil{
		fmt.Println("some Problem in getting article by url", err)
		return nil, err
	}
	if res==nil{
		page, err:=q.web_scraper_client.CallTheClient(q.ctx, request.URL)
		if err!=nil{
			fmt.Println("some Problem in getting article by url", err)
			return nil, err
		}
		// now we call the llm client 
		summary, err:=q.llm_client.GenerateDescription(q.ctx, page.Text)
		if err!=nil{
			fmt.Println("some Problem in getting article by url", err)
			return nil, err
		}

		fmt.Println("PAGE", page)
		res, err = q.articleRepository.CreateArticle(q.ctx, &models.Article{
			ArticleID:  uuid.New().String(), // generate unique ID
			Title:      page.Title,
			Link:       page.Url,
			Content:    summary,
			PubDate:    time.Now().Format(time.RFC3339),
			PubDateTZ:  "UTC",
			Language:   "en",       // default or detect
			Country:    []string{}, // empty slice
			Category:   []string{}, // empty slice
		})
		if err!=nil{
			fmt.Println("some Problem in getting article by url", err)
			return nil, err
		}
		fmt.Println("ARTICLE", res)

	}
	return res, nil
}