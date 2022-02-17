package main

import (
	"log"
	"net/http"

	"github.com/menzerath/aachen-verkehr-exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Print("starting exporter on port 9090")

	prometheus.MustRegister(exporter.NewExporter())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9090", nil))
}
