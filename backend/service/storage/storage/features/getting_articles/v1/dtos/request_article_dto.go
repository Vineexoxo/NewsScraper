package dtos



type RequestArticleDto struct {
	URL string `json:"url" validate:"required"`
}

type ResponseArticleDto struct {
	URL string 
	DESC string
	Date string
}	
