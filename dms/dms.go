package dms

import (
	"fmt"
	"math"
	"strconv"
)

// DecimalDegrees represent latitude/longitude of geographic coordinates as decimal frations of a degree.
type DecimalDegrees struct {
	Latitude  float64
	Longitude float64
}

// DMSAngle represents a single angle for degrees, miutes, seconds measurements
type DMSAngle struct {
	Degrees   int
	Minutes   int
	Seconds   float64
	Direction string
}

func (d DMSAngle) String() string {
	return fmt.Sprintf(`%d°%d'%v" %s`, d.Degrees, d.Minutes, d.Seconds, d.Direction)
}

// DMS coordinate
type DMS struct {
	Latitude  DMSAngle
	Longitude DMSAngle
}

func (d DMS) String() string {
	return fmt.Sprintf(`%s %s`, d.Latitude, d.Longitude)
}

func newDMSAngle(decimalDegreeAngle float64, direction string) DMSAngle {
	decimalDegreeAngle = math.Abs(decimalDegreeAngle)
	degrees := uint8(decimalDegreeAngle)
	minutes := uint8((decimalDegreeAngle - float64(degrees)) * 60)
	seconds := (decimalDegreeAngle - float64(degrees) - float64(minutes)/60) * 3600

	return DMSAngle{
		Degrees:   int(degrees),
		Minutes:   int(minutes),
		Seconds:   seconds,
		Direction: direction,
	}
}

// New generates a DMS position from a set of decimal degree coordinates (latitude/longitude)
func NewDMS(latlon DecimalDegrees) (*DMS, error) {
	lat, lon := latlon.Latitude, latlon.Longitude

	var latDirection, lonDirection string
	if lat > 0 {
		latDirection = "N"
	} else {
		latDirection = "S"
	}

	if lon > 0 {
		lonDirection = "E"
	} else {
		lonDirection = "W"
	}

	if lat < -90 || lat > 90 {
		return nil, fmt.Errorf("latitude must be in range of -90 and 90, found %f", lat)
	}

	if lon < -180 || lon > 180 {
		return nil, fmt.Errorf("longitude must be in range of -180 and 180, found %f", lon)
	}

	latitude := newDMSAngle(lat, latDirection)
	longitude := newDMSAngle(lon, lonDirection)

	return &DMS{Latitude: latitude, Longitude: longitude}, nil
}

func (d DMS) AutostarLongitude(longAngle DMSAngle) string {
	var autostarLongDegrees int
	var autostarLongMinutes int

	if longAngle.Direction == "W" {
		autostarLongDegrees = 360 - longAngle.Degrees
		if longAngle.Minutes == 0 {
			autostarLongMinutes = 0
		} else {
			autostarLongMinutes = 60 - longAngle.Minutes
		}

	} else {
		autostarLongDegrees = longAngle.Degrees
		autostarLongMinutes = longAngle.Minutes
	}

	autostarLongDegreesPadded := PadLeft(strconv.Itoa(autostarLongDegrees), 3) // fmt.Printf("%03s", strconv.Itoa(autostarLongDegrees))
	autostarLongitude := autostarLongDegreesPadded
	autostarLongitude += "*"
	autostarLongitude += PadLeft(strconv.Itoa(autostarLongMinutes), 2)
	return autostarLongitude
}

func (d DMS) AutostarLatitude(latAngle DMSAngle) string {
	var autostarLatDegrees int
	var autostarLatMinutes int
	if latAngle.Direction == "N" {
		autostarLatDegrees = latAngle.Degrees
		autostarLatMinutes = latAngle.Minutes
	} else {

		autostarLatDegrees = -latAngle.Degrees
		autostarLatMinutes = latAngle.Minutes
	}

	autostarLatDegreesPadded := PadLeft(strconv.Itoa(autostarLatDegrees), 2)
	autostarLatitude := autostarLatDegreesPadded
	autostarLatitude += "*"
	autostarLatitude += PadLeft(strconv.Itoa(autostarLatMinutes), 2)

	return autostarLatitude
}

func PadLeft(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}
