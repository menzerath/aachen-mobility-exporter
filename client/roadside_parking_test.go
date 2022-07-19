package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParkingData(t *testing.T) {
	data, err := GetRoadsideParkingData()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}
