package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/menzerath/aachen-verkehr-exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// DefaultPort will be used if not overridden by using the PORT environment variable.
const DefaultPort = "9090"

func main() {
	port := DefaultPort
	if value, exists := os.LookupEnv("PORT"); exists {
		port = value
	}
	log.Infof("starting exporter on port %s", port)

	prometheus.MustRegister(exporter.NewExporter())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
