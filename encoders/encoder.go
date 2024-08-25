package encoders

import (
	"math"
	"time"
)

// RADecToAltAz converts right ascension and declination to altitude and azimuth
func RADecToAltAz(ra, dec, lat, lon float64, t time.Time) (alt, az float64) {
	// Convert degrees to radians
	raRad := ra * math.Pi / 180
	decRad := dec * math.Pi / 180
	latRad := lat * math.Pi / 180

	// Calculate Local Sidereal Time (LST)
	lst := calculateLST(lon, t)

	// Calculate hour angle
	ha := lst - raRad

	// Calculate altitude
	sinAlt := math.Sin(decRad)*math.Sin(latRad) + math.Cos(decRad)*math.Cos(latRad)*math.Cos(ha)
	alt = math.Asin(sinAlt)

	// Calculate azimuth
	cosAz := (math.Sin(decRad) - math.Sin(alt)*math.Sin(latRad)) / (math.Cos(alt) * math.Cos(latRad))
	az = math.Acos(cosAz)

	// Convert radians back to degrees
	alt = alt * 180 / math.Pi
	az = az * 180 / math.Pi

	// Adjust azimuth to be measured from North
	if math.Sin(ha) > 0 {
		az = 360 - az
	}

	return alt, az
}

// AltAzToRADec converts altitude and azimuth to right ascension and declination
func AltAzToRADec(alt, az, lat, lon float64, t time.Time) (ra, dec float64) {
	// Convert degrees to radians
	altRad := alt * math.Pi / 180
	azRad := az * math.Pi / 180
	latRad := lat * math.Pi / 180

	// Calculate declination
	sinDec := math.Sin(altRad)*math.Sin(latRad) + math.Cos(altRad)*math.Cos(latRad)*math.Cos(azRad)
	dec = math.Asin(sinDec)

	// Calculate hour angle
	cosHA := (math.Sin(altRad) - math.Sin(dec)*math.Sin(latRad)) / (math.Cos(dec) * math.Cos(latRad))
	ha := math.Acos(cosHA)

	// Adjust hour angle based on azimuth
	if math.Sin(azRad) > 0 {
		ha = 2*math.Pi - ha
	}

	// Calculate Local Sidereal Time (LST)
	lst := calculateLST(lon, t)

	// Calculate right ascension
	ra = lst - ha

	// Normalize right ascension to [0, 2π)
	ra = math.Mod(ra+2*math.Pi, 2*math.Pi)

	// Convert radians back to degrees
	dec = dec * 180 / math.Pi
	ra = ra * 180 / math.Pi

	return ra, dec
}

// calculateLST calculates Local Sidereal Time
func calculateLST(lon float64, t time.Time) float64 {
	// This is a simplified LST calculation
	// For more accurate results, you may want to use a more comprehensive algorithm
	utc := t.UTC()
	d := float64(utc.YearDay()) + float64(utc.Hour())/24.0
	lst := 100.46 + 0.985647*d + lon + 15*float64(utc.Hour())
	return math.Mod(lst*math.Pi/180, 2*math.Pi)

}
