package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/menzerath/aachen-verkehr-exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// DefaultPort will be used if not overridden by using the PORT environment variable.
const DefaultPort = "9090"

func main() {
	initLogging()

	port := DefaultPort
	if value, exists := os.LookupEnv("PORT"); exists {
		port = value
	}
	slog.Info("starting exporter", slog.String("port", port))

	prometheus.MustRegister(exporter.NewExporter())
	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		slog.Error("http listen", slog.Any("error", err))
		os.Exit(1)
	}
}

func initLogging() {
	var handler slog.Handler
	if os.Getenv("MODE") == "production" {
		handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	}
	slog.SetDefault(slog.New(handler))
}
