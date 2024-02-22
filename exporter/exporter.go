package exporter

import (
	"fmt"
	"log/slog"

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
				"aachen_mobility_parking_roadside_total",
				"How many roadside parking spots are available in total.",
				[]string{"location_id", "name", "type", "sub_locations", "parent_locations", "latitude", "longitude"},
				nil,
			),
			prometheus.NewDesc(
				"aachen_mobility_parking_roadside_occupied",
				"How many roadside parking spots are occupied.",
				[]string{"location_id", "name", "type", "sub_locations", "parent_locations", "latitude", "longitude"},
				nil,
			),
			prometheus.NewDesc(
				"aachen_mobility_parking_roadside_free",
				"How many roadside parking spots are free.",
				[]string{"location_id", "name", "type", "sub_locations", "parent_locations", "latitude", "longitude"},
				nil,
			),
			prometheus.NewDesc(
				"aachen_mobility_parking_carpark_load",
				"Percentage of how many car-park parking spots are taken.",
				[]string{"id", "name", "description", "status", "trend", "latitude", "longitude"},
				nil,
			),
			prometheus.NewDesc(
				"aachen_mobility_parking_carpark_free",
				"How many car-park parking spots are free.",
				[]string{"id", "name", "description", "status", "trend", "latitude", "longitude"},
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
	roadsideParkingData, err := client.GetRoadsideParkingData()
	if err != nil {
		slog.Error("fetching roadside-parking data failed", slog.Any("error", err))
		for _, desc := range e.descriptions {
			c <- prometheus.NewInvalidMetric(desc, err)
		}
		return
	}

	carParkParkingData, err := client.GetCarParkParkingData()
	if err != nil {
		slog.Error("fetching car-park-parking data failed", slog.Any("error", err))
		for _, desc := range e.descriptions {
			c <- prometheus.NewInvalidMetric(desc, err)
		}
		return
	}

	for _, parking := range roadsideParkingData {
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

	for _, parking := range carParkParkingData.Value {
		var (
			status string
			trend  string
		)
		if parking.MultiDatastreams[0].Observations[0].Result[7] != nil {
			status = parking.MultiDatastreams[0].Observations[0].Result[7].(string)
		}
		if parking.MultiDatastreams[0].Observations[0].Result[4] != nil {
			trend = parking.MultiDatastreams[0].Observations[0].Result[4].(string)
		}

		c <- prometheus.MustNewConstMetric(
			e.descriptions[3],
			prometheus.GaugeValue,
			parking.MultiDatastreams[0].Observations[0].Result[3].(float64),
			fmt.Sprintf("%d", parking.ID),
			parking.Name,
			parking.Description,
			status,
			trend,
			fmt.Sprintf("%f", parking.Locations[0].Location.Geometry.Coordinates[1]),
			fmt.Sprintf("%f", parking.Locations[0].Location.Geometry.Coordinates[0]),
		)
		c <- prometheus.MustNewConstMetric(
			e.descriptions[4],
			prometheus.GaugeValue,
			parking.MultiDatastreams[0].Observations[0].Result[0].(float64),
			fmt.Sprintf("%d", parking.ID),
			parking.Name,
			parking.Description,
			status,
			trend,
			fmt.Sprintf("%f", parking.Locations[0].Location.Geometry.Coordinates[1]),
			fmt.Sprintf("%f", parking.Locations[0].Location.Geometry.Coordinates[0]),
		)
	}
}
