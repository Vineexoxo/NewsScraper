package main

import (
	"go.uber.org/fx"
	"/http/http_client"
)


func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				httpclient.NewHttpClient,
			),
			fx.Invoke(fetchNews),
		),
	).Run()	
	fmt.Println("News fetching service started")
}



