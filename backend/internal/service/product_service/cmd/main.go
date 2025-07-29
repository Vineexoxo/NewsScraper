package main

import (
	"go.uber.org/fx"
	"github.com/shishir54234/NewsScraper/backend/internal/pkg/http/httpclient"
)


func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				httpclient.NewHttpClient,
			),
			fx.Invoke(fetchNews),



		)


	)
}



