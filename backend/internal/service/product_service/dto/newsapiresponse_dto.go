package dto

type NewsAPIResponse struct {
	Status string `json:"status"`
	TotalResults int `json:"totalResults"`
	Results []Article `json:"results"`
	NextPage string `json:"nextPage"`
}