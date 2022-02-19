package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const parkingURL = "https://verkehr.aachen.de/api/sonah/api/v2/locations"

// Parking describes the current parking situation at a specific location.
type Parking struct {
	LocationID int64  `json:"LocationID"`
	Name       string `json:"Name"`
	Type       string `json:"Type"`

	TotalParking    float64 `json:"TotalParking"`
	OccupiedParking float64 `json:"OccupiedParking"`
	FreeParking     float64 `json:"FreeParking"`

	SubLocations    []string `json:"SubLocations"`
	ParentLocations []string `json:"ParentLocations"`

	Positions ParkingPosition `json:"Positions"`
}

// ParkingPosition contains two locations of a parking position.
type ParkingPosition struct {
	Center     ParkingPositionCoordinates `json:"Center"`
	Navigation ParkingPositionCoordinates `json:"Navigation"`
}

// ParkingPositionCoordinates contains the coordinates of a parking position.
type ParkingPositionCoordinates struct {
	Lat  float64 `json:"Lat"`
	Long float64 `json:"Long"`
}

// GetParkingData returns the current parking situation from the public api.
func GetParkingData() ([]Parking, error) {
	response, err := http.Get(parkingURL)
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

	var data []Parking
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
