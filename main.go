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

	"github.com/jlcoulter/ipam/internal/config"
	"github.com/jlcoulter/ipam/internal/handler"
	"github.com/jlcoulter/ipam/internal/logging"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "ipam",
		Short: "IP address management tracker",
	}

	rootCmd.AddCommand(
		allocateCmd,
		releaseCmd,
		listCmd,
		showCmd,
		subnetCmd,
		serveCmd,
		scanCmd,
		importCmd,
	)
	rootCmd.Execute()

	// TODO Review whats required from below template

	cfg := config.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: parseLogLevel(cfg.LogLevel),
	}))
	slog.SetDefault(logger)

	r := chi.NewRouter()
	r.Use(logging.RequestLogger(logger))

	r.Get("/healthz", handler.Healthz)
	r.Get("/readyz", handler.Readyz)

	// Example routes — delete and replace with your own
	r.Route("/api", func(r chi.Router) {
		r.Get("/example", handler.ListExamples)
		r.Post("/example", handler.CreateExample)
		r.Get("/example/{id}", handler.GetExample)
		r.Delete("/example/{id}", handler.DeleteExample)
	})

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

