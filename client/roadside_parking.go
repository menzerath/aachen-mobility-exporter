package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const roadsideParkingURL = "https://verkehr.aachen.de/api/sonah/api/v2/locations"

// RoadsideParking describes the current roadside-parking situation at a specific location.
type RoadsideParking struct {
	LocationID int64  `json:"LocationID"`
	Name       string `json:"Name"`
	Type       string `json:"Type"`

	TotalParking    float64 `json:"TotalParking"`
	OccupiedParking float64 `json:"OccupiedParking"`
	FreeParking     float64 `json:"FreeParking"`

	SubLocations    []string `json:"SubLocations"`
	ParentLocations []string `json:"ParentLocations"`

	Positions RoadsideParkingPosition `json:"Positions"`
}

// RoadsideParkingPosition contains two locations of a roadside-parking position.
type RoadsideParkingPosition struct {
	Center     RoadsideParkingPositionCoordinates `json:"Center"`
	Navigation RoadsideParkingPositionCoordinates `json:"Navigation"`
}

// RoadsideParkingPositionCoordinates contains the coordinates of a roadside-parking position.
type RoadsideParkingPositionCoordinates struct {
	Lat  float64 `json:"Lat"`
	Long float64 `json:"Long"`
}

// GetRoadsideParkingData returns the current roadside-parking situation from the public api.
func GetRoadsideParkingData() ([]RoadsideParking, error) {
	response, err := http.Get(roadsideParkingURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("http error %d", response.StatusCode)
		}
		return nil, fmt.Errorf("http error %d: %s", response.StatusCode, message)
	}

	var data []RoadsideParking
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
