package dtos

type GenerateDescriptionRequest struct {
	URL        string `json:"url"`
	Description string `json:"description"`
}

type GenerateDescriptionResponse struct {
	URL        string `json:"url"`
	Description string `json:"description"`
}