package dtos



type RequestArticleDto struct {
	URL string `json:"url" validate:"required"`
}

type ResponseRequestDto struct {
	URL string 
	DESC string
	Date string
}	
