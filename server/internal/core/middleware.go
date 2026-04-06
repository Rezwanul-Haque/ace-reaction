package core

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(e *echo.Echo, cfg *Config) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	origins := []string{"*"}
	if cfg.AllowedOrigin != "" {
		origins = []string{cfg.AllowedOrigin}
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
}
