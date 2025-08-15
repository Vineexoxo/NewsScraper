package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	// "reflect"

	// "database/sql"
	// "fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"

	// "github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	pgsql "github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/pkg/utils"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/contracts"
	// "gorm.io/gorm"
)
func fromNull(ns sql.NullString) string {
    if ns.Valid {
        return ns.String
    }
    return ""
}
type PostgresArticleRepository struct {
	log  logger.ILogger
	cfg  *pgsql.PostgresConfig
	db   *pgxpool.Pool
	gorm *pgsql.PostgresDB
}

func NewPostgresArticleRepository(log logger.ILogger, cfg *pgsql.PostgresConfig, gorm *pgsql.PostgresDB) contracts.ArticleRepository {
	return &PostgresArticleRepository{log: log, cfg: cfg, gorm: gorm}
}

func (p *PostgresArticleRepository) GetAllArticles(ctx context.Context, listQuery *utils.ListQuery) (
    *utils.ListResult[*models.Article], error) {
        fmt.Println("NITHEZZ IS A FUCKING NIGGER")
    stats := p.gorm.DB.Stats()
    fmt.Printf(`
    GORM DB Connection Pool Stats:
    ------------------------------------
    Max Open Connections : %d
    Open Connections     : %d
    In Use Connections   : %d
    Idle Connections     : %d
    Wait Count           : %d
    Wait Duration        : %s
    Max Idle Closed      : %d
    Max Idle Time Closed : %d
    Max Lifetime Closed  : %d
    `,
        stats.MaxOpenConnections,
        stats.OpenConnections,
        stats.InUse,
        stats.Idle,
        stats.WaitCount,
        stats.WaitDuration,
        stats.MaxIdleClosed,
        stats.MaxIdleTimeClosed,
        stats.MaxLifetimeClosed,
    )

    // ---- Build SQL ----
    baseQuery := "SELECT * FROM articles"
    whereParts := []string{}
    args := []interface{}{}
    argIndex := 1 // for $1, $2, etc.

    comparisonMap := map[string]string{
        "eq":  "=",
        "neq": "!=",
        "gt":  ">",
        "lt":  "<",
        "gte": ">=",
        "lte": "<=",
        "like": "ILIKE",
    }

    // Apply filters
    for _, f := range listQuery.Filters {
        if f.Field == "" || f.Value == "" {
            continue
        }
        if op, ok := comparisonMap[strings.ToLower(f.Comparison)]; ok {
            if op == "ILIKE" {
                whereParts = append(whereParts, fmt.Sprintf("%s %s $%d", f.Field, op, argIndex))
                args = append(args, "%"+f.Value+"%")
            } else {
                whereParts = append(whereParts, fmt.Sprintf("%s %s $%d", f.Field, op, argIndex))
                args = append(args, f.Value)
            }
            argIndex++
        }
    }

    if len(whereParts) > 0 {
        baseQuery += " WHERE " + strings.Join(whereParts, " AND ")
    }

    // Apply ordering
    if listQuery.OrderBy != "" {
        // ⚠️ sanitize allowed columns if needed to prevent SQL injection
        baseQuery += " ORDER BY " + listQuery.OrderBy
    }

    // Pagination
    limit := listQuery.Size
    if limit <= 0 {
        limit = 10
    }
    offset := (listQuery.Page - 1) * limit
    if offset < 0 {
        offset = 0
    }

    baseQuery += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

    // ---- Execute SQL ----
    rows, err := p.gorm.DB.Query(baseQuery, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var articles []*models.Article
    for rows.Next() {
        var a models.Article
        var (
            imageURL, videoURL sql.NullString
            keywordsArr        pq.StringArray
            creatorArr         pq.StringArray
            countryArr         pq.StringArray
            categoryArr        pq.StringArray
        )

        err := rows.Scan(
            &a.ArticleID,
            &a.Title,
            &a.Link,
            &keywordsArr,
            &creatorArr,
            &a.Description,
            &a.Content,
            &a.PubDate,
            &a.PubDateTZ,
            &imageURL,
            &videoURL,
            &a.SourceID,
            &a.SourceName,
            &a.SourcePriority,
            &a.SourceURL,
            &a.SourceIcon,
            &a.Language,
            &countryArr,
            &categoryArr,
            &a.Sentiment,
            &a.SentimentStats,
            &a.AITag,
            &a.AIRegion,
            &a.AIOrg,
            &a.AISummary,
            &a.AIContent,
            &a.Duplicate,
        )
        if err != nil {
            return nil, err
        }

        a.ImageURL = fromNull(imageURL)
        a.VideoURL = fromNull(videoURL)
        a.Keywords = []string(keywordsArr)
        a.Creator = []string(creatorArr)
        a.Country = []string(countryArr)
        a.Category = []string(categoryArr)

        articles = append(articles, &a)
    }








    return &utils.ListResult[*models.Article]{Items: articles}, nil
}
func (p *PostgresArticleRepository) CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error) {
	query := `
		INSERT INTO articles (
			article_id, title, link, keywords, creator, description, content, pub_date, pub_date_tz,
			image_url, video_url, source_id, source_name, source_priority, source_url, source_icon,
			language, country, category, sentiment, sentiment_stats, ai_tag, ai_region, ai_org,
			ai_summary, ai_content, duplicate
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9,
			$10, $11, $12, $13, $14, $15, $16,
			$17, $18, $19, $20, $21, $22, $23, $24,
			$25, $26, $27
		)
	`

	_, err := p.gorm.DB.Query(query,
		article.ArticleID,
		article.Title,
		article.Link,
		pq.Array(article.Keywords),
		pq.Array(article.Creator),
		article.Description,
		article.Content,
		article.PubDate,
		article.PubDateTZ,
		article.ImageURL,
		article.VideoURL,
		article.SourceID,
		article.SourceName,
		article.SourcePriority,
		article.SourceURL,
		article.SourceIcon,
		article.Language,
		pq.Array(article.Country),
		pq.Array(article.Category),
		article.Sentiment,
		article.SentimentStats,
		article.AITag,
		article.AIRegion,
		article.AIOrg,
		article.AISummary,
		article.AIContent,
		article.Duplicate,
	)

	if err != nil {
		return nil, err
	}

	return article, nil
}
func (p *PostgresArticleRepository) SearchArticles(ctx context.Context, searchText string, listQuery *utils.ListQuery) (*utils.ListResult[*models.Article], error) {

	// whereQuery := fmt.Sprintf("%s IN (?)", "Name")
	// query := p.gorm.Where(whereQuery, searchText)

	// result, err := pgsql.Paginate[*models.Article](ctx, listQuery, query)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return nil, nil
}


func (p *PostgresArticleRepository) GetArticleByUrl(ctx context.Context, 
url string) (*models.Article, error) {
    baseQuery := "SELECT * FROM articles where link = $1"
    rows, err := p.gorm.DB.Query(baseQuery, url)
    if err != nil {
        return nil, err
    }
    var article *models.Article = nil
    for rows.Next() {
        var a models.Article
        var (
            imageURL, videoURL sql.NullString
            keywordsArr        pq.StringArray
            creatorArr         pq.StringArray
            countryArr         pq.StringArray
            categoryArr        pq.StringArray
        )

        err := rows.Scan(
            &a.ArticleID,
            &a.Title,
            &a.Link,
            &keywordsArr,
            &creatorArr,
            &a.Description,
            &a.Content,
            &a.PubDate,
            &a.PubDateTZ,
            &imageURL,
            &videoURL,
            &a.SourceID,
            &a.SourceName,
            &a.SourcePriority,
            &a.SourceURL,
            &a.SourceIcon,
            &a.Language,
            &countryArr,
            &categoryArr,
            &a.Sentiment,
            &a.SentimentStats,
            &a.AITag,
            &a.AIRegion,
            &a.AIOrg,
            &a.AISummary,
            &a.AIContent,
            &a.Duplicate,
        )
        if err != nil {
            return nil, err
        }

        a.ImageURL = fromNull(imageURL)
        a.VideoURL = fromNull(videoURL)
        a.Keywords = []string(keywordsArr)
        a.Creator = []string(creatorArr)
        a.Country = []string(countryArr)
        a.Category = []string(categoryArr)

        article = &a
    }
	return article, nil
}


func (p *PostgresArticleRepository) UpdateArticle(ctx context.Context, updateArticle *models.Article) (*models.Article, error) {

	// if err := p.gorm.Save(updateArticle).Error; err != nil {
	// 	return nil, errors.Wrap(err, fmt.Sprintf("error in updating article with id %s into the database."))
	// }

	// return updateArticle, nil
	return nil, nil
}

func (p *PostgresArticleRepository) DeleteArticleByID(ctx context.Context, uuid uuid.UUID) error {

	// var article models.Article

	// if err := p.gorm.First(article, uuid).Error; err != nil {
	// 	return errors.Wrap(err, fmt.Sprintf("can't find the article with id %s into the database.", uuid))
	// }

	// if err := p.gorm.Delete(article).Error; err != nil {
	// 	return errors.Wrap(err, "error in the deleting article into the database.")
	// }

	return nil
}
