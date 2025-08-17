package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	echoserver "github.com/shishir54234/NewsScraper/backend/pkg/http/echo/server"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/service/storage/config"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, log logger.ILogger, e *echo.Echo, 
ctx context.Context, cfg *config.Config) error {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // your React dev server
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			
			go func() {
				if err := echoserver.RunHttpServer(ctx, 
				e, log, cfg.Echo); !errors.Is(err, http.ErrServerClosed) {
					fmt.Printf("error running http server: %v", err)
				}
			}()

			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, config.GetMicroserviceName(cfg.ServiceName))
			})

			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Infof("all servers shutdown gracefully...")
			return nil
		},
	})

	return nil
}
