package dms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDMS(t *testing.T) {
	dms, err := NewDMS(DecimalDegrees{Latitude: 23.33, Longitude: 42.55})
	assert.Nil(t, err)

	// test latitude
	assert.Equal(t, dms.Latitude.Degrees, 23)
	assert.Equal(t, dms.Latitude.Minutes, 19)
	assert.Equal(t, dms.Latitude.Seconds, 47.99999999999392)
	assert.Equal(t, dms.Latitude.Direction, "N")
	assert.Equal(t, dms.Latitude.String(), `23°19'47.99999999999392" N`)

	// test longitude
	assert.Equal(t, dms.Longitude.Degrees, 42)
	assert.Equal(t, dms.Longitude.Minutes, 32)
	assert.Equal(t, dms.Longitude.Seconds, 59.9999999999898)
	assert.Equal(t, dms.Longitude.Direction, "E")
	assert.Equal(t, dms.Longitude.String(), `42°32'59.9999999999898" E`)

	assert.Equal(t, dms.String(), `23°19'47.99999999999392" N 42°32'59.9999999999898" E`)

	dms, err = NewDMS(DecimalDegrees{Latitude: -66.434323, Longitude: -115.25})
	assert.Nil(t, err)

	// test latitude
	assert.Equal(t, dms.Latitude.Degrees, 66)
	assert.Equal(t, dms.Latitude.Minutes, 26)
	assert.Equal(t, dms.Latitude.Seconds, 3.5628000000223814)
	assert.Equal(t, dms.Latitude.Direction, "S")
	assert.Equal(t, dms.Latitude.String(), `66°26'3.5628000000223814" S`)

	// test longitude
	assert.Equal(t, dms.Longitude.Degrees, 115)
	assert.Equal(t, dms.Longitude.Minutes, 15)
	assert.Equal(t, dms.Longitude.Seconds, 0.0)
	assert.Equal(t, dms.Longitude.Direction, "W")
	assert.Equal(t, dms.Longitude.String(), `115°15'0" W`)

	assert.Equal(t, dms.String(), `66°26'3.5628000000223814" S 115°15'0" W`)
}

func TestAutostarLongitude(t *testing.T) {
	dms, _ := NewDMS(DecimalDegrees{Latitude: 41.82326, Longitude: 12.44474})

	// Test east longitude
	autostarLong := dms.AutostarLongitude(dms.Longitude)
	assert.Equal(t, "012*26", autostarLong)

	// Test west longitude
	dms.Longitude.Direction = "W"
	autostarLong = dms.AutostarLongitude(dms.Longitude)
	assert.Equal(t, "348*34", autostarLong)

	// Test 0 degree longitude
	dms.Longitude = DMSAngle{Degrees: 0, Minutes: 0, Seconds: 0.0, Direction: "E"}
	autostarLong = dms.AutostarLongitude(dms.Longitude)
	assert.Equal(t, "000*00", autostarLong)

	// Test 180 degree longitude
	dms.Longitude = DMSAngle{Degrees: 180, Minutes: 0, Seconds: 0.0, Direction: "W"}
	autostarLong = dms.AutostarLongitude(dms.Longitude)
	assert.Equal(t, "180*00", autostarLong)

	// Test 90 E degree longitude
	dms.Longitude = DMSAngle{Degrees: 90, Minutes: 0, Seconds: 0.0, Direction: "E"}
	autostarLong = dms.AutostarLongitude(dms.Longitude)
	assert.Equal(t, "090*00", autostarLong)

	// Test 90 W degree longitude
	dms.Longitude = DMSAngle{Degrees: 90, Minutes: 0, Seconds: 0.0, Direction: "W"}
	autostarLong = dms.AutostarLongitude(dms.Longitude)
	assert.Equal(t, "270*00", autostarLong)

	dms.Longitude = DMSAngle{Degrees: 180, Minutes: 0, Seconds: 0.0, Direction: "E"}
	autostarLong = dms.AutostarLongitude(dms.Longitude)
	assert.Equal(t, "180*00", autostarLong)
}

func TestAutostarLatitude(t *testing.T) {
	dms, _ := NewDMS(DecimalDegrees{Latitude: 41.82326, Longitude: 12.44474})

	// Test north latitude
	autostarLat := dms.AutostarLatitude(dms.Latitude)
	assert.Equal(t, "41*49", autostarLat)

	// Test south latitude
	dms.Latitude.Direction = "S"
	autostarLat = dms.AutostarLatitude(dms.Latitude)
	assert.Equal(t, "-41*49", autostarLat)

	// Test 0 degree latitude
	dms.Latitude = DMSAngle{Degrees: 0, Minutes: 0, Seconds: 0.0, Direction: "N"}
	autostarLat = dms.AutostarLatitude(dms.Latitude)
	assert.Equal(t, "00*00", autostarLat)

	// Test 90 degree latitude
	dms.Latitude = DMSAngle{Degrees: 90, Minutes: 0, Seconds: 0.0, Direction: "N"}
	autostarLat = dms.AutostarLatitude(dms.Latitude)
	assert.Equal(t, "90*00", autostarLat)

	// Test -90 degree latitude
	dms.Latitude = DMSAngle{Degrees: 90, Minutes: 0, Seconds: 0.0, Direction: "S"}
	autostarLat = dms.AutostarLatitude(dms.Latitude)
	assert.Equal(t, "-90*00", autostarLat)

}
