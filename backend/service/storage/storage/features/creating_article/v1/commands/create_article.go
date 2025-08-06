package commands

import "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/creating_article/v1/dtos"

type CreateArticle struct {
	ArticleID      string   `json:"article_id"`
	Title          string   `json:"title"`
	Link           string   `json:"link"`
	Keywords       []string `json:"keywords"`
	Creator        []string `json:"creator"` // nullable
	Description    string   `json:"description"`
	Content        string   `json:"content"`
	PubDate        string   `json:"pubDate"`
	PubDateTZ      string   `json:"pubDateTZ"`
	ImageURL       string   `json:"image_url"` // nullable
	VideoURL       string   `json:"video_url"` // nullable
	SourceID       string   `json:"source_id"`
	SourceName     string   `json:"source_name"`
	SourcePriority int      `json:"source_priority"`
	SourceURL      string   `json:"source_url"`
	SourceIcon     string   `json:"source_icon"`
	Language       string   `json:"language"`
	Country        []string `json:"country"`
	Category       []string `json:"category"`
	Sentiment      string   `json:"sentiment"`
	SentimentStats string   `json:"sentiment_stats"`
	AITag          string   `json:"ai_tag"`
	AIRegion       string   `json:"ai_region"`
	AIOrg          string   `json:"ai_org"`
	AISummary      string   `json:"ai_summary"`
	AIContent      string   `json:"ai_content"`
	Duplicate      bool     `json:"duplicate"`
}

func NewCreateArticle(dto dtos.CreateArticleRequestDto) *CreateArticle {
	return &CreateArticle{
		ArticleID:      dto.ArticleID,
		Title:          dto.Title,
		Link:           dto.Link,
		Keywords:       dto.Keywords,
		Creator:        dto.Creator,
		Description:    dto.Description,
		Content:        dto.Content,
		PubDate:        dto.PubDate,
		PubDateTZ:      dto.PubDateTZ,
		ImageURL:       dto.ImageURL,
		VideoURL:       dto.VideoURL,
		SourceID:       dto.SourceID,
		SourceName:     dto.SourceName,
		SourcePriority: dto.SourcePriority,
		SourceURL:      dto.SourceURL,
		SourceIcon:     dto.SourceIcon,
		Language:       dto.Language,
		Country:        dto.Country,
		Category:       dto.Category,
		Sentiment:      dto.Sentiment,
		SentimentStats: dto.SentimentStats,
		AITag:          dto.AITag,
		AIRegion:       dto.AIRegion,
		AIOrg:          dto.AIOrg,
		AISummary:      dto.AISummary,
		AIContent:      dto.AIContent,
		Duplicate:      dto.Duplicate,
	}
}
