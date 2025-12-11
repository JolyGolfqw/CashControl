package main

import (
	"log/slog"
	"os"

	"cashcontrol/internal/config"
	"cashcontrol/internal/database"
	"cashcontrol/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger := initLogger()

	if err := database.Init(cfg); err != nil {
		logger.Error("failed to init database", slog.String("error", err.Error()))
		panic(err)
	}

	if err := database.Migrate(); err != nil {
		logger.Error("failed to migrate database", slog.String("error", err.Error()))
		panic(err)
	}

	defer func() {
		if err := database.Close(); err != nil {
			logger.Error("failed to close database", slog.String("error", err.Error()))
		}
	}()

	// Инициализация маршрутизатора
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Сервер работает",
			"status":  "ok",
		})
	})

	handlers.RegisterRoutes(router, database.DB, logger)

	logger.Info("starting server", slog.String("address", cfg.ServerAddress))

	if err := router.Run(cfg.ServerAddress); err != nil {
		logger.Error("failed to start server", slog.String("error", err.Error()))
		panic(err)
	}
}

// initLogger инициализирует структурированный логгер
func initLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler)
}
