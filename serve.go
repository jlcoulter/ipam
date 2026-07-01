package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"

	"github.com/jlcoulter/ipam/internal/api"
	"github.com/jlcoulter/ipam/internal/config"
	"github.com/jlcoulter/ipam/internal/logging"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start REST API server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()

		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: parseLogLevel(cfg.LogLevel),
		}))
		slog.SetDefault(logger)

		r := chi.NewRouter()
		r.Use(logging.RequestLogger(logger))

		r.Get("/healthz", api.Healthz)
		r.Get("/readyz", api.Readyz)

		// TODO: Register IPAM API routes

		srv := &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      r,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		}

		go func() {
			slog.Info("starting server", "port", cfg.Port)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				slog.Error("server error", "error", err)
				os.Exit(1)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		slog.Info("shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("server forced to shutdown", "error", err)
		}
		slog.Info("server stopped")
	},
}

func init() {
	serveCmd.Flags().StringP("addr", "a", ":8080", "listen address")
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
