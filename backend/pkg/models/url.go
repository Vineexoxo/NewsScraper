package models

import "time"

type URLStatus string
const (
	URLStatusPending   URLStatus = "pending"
	URLStatusScraped   URLStatus = "scraped"
	URLStatusFailed    URLStatus = "failed"
	URLStatusArchived  URLStatus = "archived"
	URLStatusProcessing URLStatus = "processing"
)


type URL struct {
	ID            string    `json:"id"`
	URL           string    `json:"url"`
	Title         string    `json:"title"`
	Domain        string    `json:"domain"`
	Category      string    `json:"category"`
	Status        URLStatus `json:"status"`
	CreatedAt     string    `json:"created_at"`
	LastScrapedAt *time.Time `json:"last_scraped_at,omitempty"`
}

type URLBatch struct{
	URLs []string `json:"urls" validate:"required,min=1,dive,url"`
	Category string `json:"category" validate:"required"`
}




