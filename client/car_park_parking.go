package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const carParkParkingURL = "https://verkehr-mp.aachen.de/FROST-Server/v1.1/Things?%24filter=Datastreams%2Fproperties%2Ftype%20eq%20%27Parkobjekt%27&%24expand=Locations%2CDatastreams%2FObservations(%24top%3D1%3B%24orderby%3DphenomenonTime%20desc)&%24count=true"

// CarParkParking describes the current car-park-parking situation at a specific location.
type CarParkParking struct {
	Count int64 `json:"@iot.count"`
	Value []struct {
		ID          int64  `json:"@iot.id"`
		Name        string `json:"name"`
		Description string `json:"description"`

		Locations []struct {
			Location struct {
				Geometry struct {
					Type        string    `json:"type"`
					Coordinates []float64 `json:"coordinates"`
				} `json:"geometry"`
			} `json:"location"`
		} `json:"Locations"`

		Datastreams []struct {
			Observations []struct {
				Timestamp  string `json:"resultTime"`
				Result     string `json:"result"`
				Parameters struct {
					Load  float64 `json:"load"`
					Trend string  `json:"trend"`
				} `json:"parameters"`
			} `json:"Observations"`
		} `json:"Datastreams"`
	} `json:"value"`
}

// GetCarParkParkingData returns the current car-park-parking situation from the public api.
func GetCarParkParkingData() (CarParkParking, error) {
	response, err := http.Get(carParkParkingURL)
	if err != nil {
		return CarParkParking{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)
		if err != nil {
			return CarParkParking{}, fmt.Errorf("http error %d", response.StatusCode)
		}
		return CarParkParking{}, fmt.Errorf("http error %d: %s", response.StatusCode, message)
	}

	var data CarParkParking
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
