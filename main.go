package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/menzerath/aachen-verkehr-exporter/exporter"
	"github.com/menzerath/aachen-verkehr-exporter/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

// DefaultPort will be used if not overridden by using the PORT environment variable.
const DefaultPort = "9090"

func main() {
	log := log.NewLogger()
	defer log.Sync()

	port := DefaultPort
	if value, exists := os.LookupEnv("PORT"); exists {
		port = value
	}
	log.Info("starting exporter", zap.String("port", port))

	prometheus.MustRegister(exporter.NewExporter())
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal("ListenAndServe failed", zap.Error(http.ListenAndServe(fmt.Sprintf(":%s", port), nil)))
}
