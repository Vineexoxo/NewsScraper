package mappings

import (
	"fmt"

	"github.com/shishir54234/NewsScraper/backend/pkg/mapper"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/dtos"
)

// import (
// 	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/mapper"
// 	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/product/dtos"
// 	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/product/features/creating_product/v1/events"
// 	events2 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/product/features/updating_product/v1/events"
// 	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/product/models"
// )

func ConfigureMappings() error {
	fmt.Println("atleast this got invoked")
	// err := mapper.CreateMap[*models.Product, *dtos.ProductDto]()
	// if err != nil {
	// 	return err
	// }

	// err = mapper.CreateMap[*models.Product, *events.ProductCreated]()
	// if err != nil {
	// 	return err
	// }

	// err = mapper.CreateMap[*models.Product, *events2.ProductUpdated]()
	// if err != nil {
	// 	return err
	// }
	err:= mapper.CreateCustomMap[*models.Article, *dtos.ResponseArticleDto]( 
		func(a *models.Article) *dtos.ResponseArticleDto {
		
		return &dtos.ResponseArticleDto{URL: a.Link, DESC: a.Description, Date: a.PubDate}
	})

	if err != nil {
		fmt.Println("ERRRRRRRRR", err)
		return err

	}	

	return nil
}
