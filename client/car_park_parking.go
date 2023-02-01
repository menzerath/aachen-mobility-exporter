package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const carParkParkingURL = "https://verkehr.aachen.de/api/sensorthings/Things?$count=false&$filter=properties/type%20eq%20%27ParkingRecord%27%20and%20properties/archive%20eq%20%27false%27&$expand=Locations,MultiDatastreams%2FObservations(%24top%3D1%3B%24orderby%3DphenomenonTime%20desc%3B%24select%3Dresult%2CphenomenonTime%2Cparameters),Datastreams%2FObservedProperty(%24select%3D%40iot.id%2Cname)&$top=300&$select=@iot.id,description,name,properties/props&$orderBy=name"

// CarParkParking describes the current car-park-parking situation at a specific location.
type CarParkParking struct {
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

		MultiDatastreams []struct {
			Observations []struct {
				Timestamp string `json:"phenomenonTime"`

				// 0: ParkingNumberOfVacantSpaces
				// 1: ParkingNumberOfOccupiedSpaces
				// 2: ParkingNumberOfVehicles
				// 3: ParkingOccupancy (%)
				// 4: ParkingOccupancyTrend
				// 5: NumberOfIncomingVehicles
				// 6: NumberOfOutgoingVehicles
				// 7: ParkingSiteStatus
				// 8: ParkingSiteOpeningStatus
				// 9: ParkingNumberOfSpacesOverride
				Result []any `json:"result"`
			} `json:"Observations"`
		} `json:"MultiDatastreams"`
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
