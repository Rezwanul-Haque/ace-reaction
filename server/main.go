package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rezwanul-haque/reflex-card-game/server/internal/core"
	"github.com/rezwanul-haque/reflex-card-game/server/internal/features/game"
	"github.com/rezwanul-haque/reflex-card-game/server/internal/features/health"
	"github.com/rezwanul-haque/reflex-card-game/server/internal/features/room"
)

func main() {
	cfg := core.LoadConfig()

	e := echo.New()
	e.HideBanner = true

	// Middleware
	core.SetupMiddleware(e, cfg)

	// Dependencies
	roomRepo := room.NewMemoryRoomRepository()
	roomSvc := room.NewRoomService(roomRepo)
	gameSvc := game.NewGameService()

	// Register routes
	api := e.Group("/api")
	roomHandler := room.NewRoomHandler(roomSvc)
	roomHandler.RegisterRoutes(api)

	wsHandler := game.NewWSHandler(roomSvc, gameSvc)
	wsHandler.RegisterRoutes(e)

	health.RegisterRoutes(e)

	// Serve static frontend files in production
	if staticDir := os.Getenv("STATIC_DIR"); staticDir != "" {
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:   staticDir,
			Index:  "index.html",
			HTML5:  true,
			Browse: false,
		}))
	}

	// Start server in goroutine
	go func() {
		log.Printf("Server starting on :%s", cfg.Port)
		if err := e.Start(":" + cfg.Port); err != nil {
			log.Printf("Server stopped: %v", err)
		}
	}()

	// Graceful shutdown
	gracefulShutdown(e)
}

func gracefulShutdown(e *echo.Echo) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server shutdown gracefully")
}
