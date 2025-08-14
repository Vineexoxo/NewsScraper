package models

// Message structures for RabbitMQ communication for web scraper service
type ScrapeJobMessage struct {
	JobID     string `json:"job_id"`
	URL       string `json:"url"`
	UserAgent string `json:"user_agent"`
}
