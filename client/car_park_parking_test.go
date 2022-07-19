package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCarParkParkingData(t *testing.T) {
	data, err := GetCarParkParkingData()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}
