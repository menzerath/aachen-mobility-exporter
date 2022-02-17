package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/menzerath/aachen-verkehr-exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// DefaultPort will be used if not overriden by using the PORT environment variable.
const DefaultPort = 9090

func main() {
	port := DefaultPort
	if value, exists := os.LookupEnv("PORT"); exists {
		newPort, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			log.Fatalf("invalid port: %s", err)
			return
		}
		port = int(newPort)
	}
	log.Infof("starting exporter on port %d", port)

	prometheus.MustRegister(exporter.NewExporter())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
