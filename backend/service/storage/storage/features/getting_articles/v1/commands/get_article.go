package commands

import (
	"github.com/shishir54234/NewsScraper/backend/pkg/utils"
)

type GetArticles struct {
	*utils.ListQuery
}


func NewGetArticles(qry *utils.ListQuery) *GetArticles {
	return &GetArticles{ListQuery: qry}
}
