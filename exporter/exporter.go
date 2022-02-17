package exporter

import (
	"fmt"
	"log"

	"github.com/menzerath/aachen-verkehr-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
)

// Exporter collects and exposes parking data from the public api.
type Exporter struct {
	descriptions []*prometheus.Desc
}

// NewExporter creates and returns a new Exporter.
func NewExporter() *Exporter {
	return &Exporter{
		descriptions: []*prometheus.Desc{
			prometheus.NewDesc(
				"aachen_mobility_parking_total",
				"How many parking spots are available in total.",
				[]string{"location_id", "name", "type", "sub_locations", "parent_locations", "latitude", "longitude"},
				nil,
			),
			prometheus.NewDesc(
				"aachen_mobility_parking_occupied",
				"How many parking spots are occupied.",
				[]string{"location_id", "name", "type", "sub_locations", "parent_locations", "latitude", "longitude"},
				nil,
			),
			prometheus.NewDesc(
				"aachen_mobility_parking_free",
				"How many parking spots are free.",
				[]string{"location_id", "name", "type", "sub_locations", "parent_locations", "latitude", "longitude"},
				nil,
			),
		},
	}
}

// Describe exposes all metric descriptions.
func (e *Exporter) Describe(c chan<- *prometheus.Desc) {
	for _, desc := range e.descriptions {
		c <- desc
	}
}

// Collect collects and exposes all metric values.
func (e *Exporter) Collect(c chan<- prometheus.Metric) {
	parkingData, err := client.GetParkingData()
	if err != nil {
		log.Printf("fetching parking data failed: %s", err)
		for _, desc := range e.descriptions {
			c <- prometheus.NewInvalidMetric(desc, err)
		}
		return
	}

	for _, parking := range parkingData {
		c <- prometheus.MustNewConstMetric(
			e.descriptions[0],
			prometheus.GaugeValue,
			parking.TotalParking,
			fmt.Sprintf("%d", parking.LocationID),
			parking.Name,
			parking.Type,
			fmt.Sprintf("%d", len(parking.SubLocations)),
			fmt.Sprintf("%d", len(parking.ParentLocations)),
			fmt.Sprintf("%f", parking.Positions.Center.Lat),
			fmt.Sprintf("%f", parking.Positions.Center.Long),
		)
		c <- prometheus.MustNewConstMetric(
			e.descriptions[1],
			prometheus.GaugeValue,
			parking.OccupiedParking,
			fmt.Sprintf("%d", parking.LocationID),
			parking.Name,
			parking.Type,
			fmt.Sprintf("%d", len(parking.SubLocations)),
			fmt.Sprintf("%d", len(parking.ParentLocations)),
			fmt.Sprintf("%f", parking.Positions.Center.Lat),
			fmt.Sprintf("%f", parking.Positions.Center.Long),
		)
		c <- prometheus.MustNewConstMetric(
			e.descriptions[2],
			prometheus.GaugeValue,
			parking.FreeParking,
			fmt.Sprintf("%d", parking.LocationID),
			parking.Name,
			parking.Type,
			fmt.Sprintf("%d", len(parking.SubLocations)),
			fmt.Sprintf("%d", len(parking.ParentLocations)),
			fmt.Sprintf("%f", parking.Positions.Center.Lat),
			fmt.Sprintf("%f", parking.Positions.Center.Long),
		)
	}
}
