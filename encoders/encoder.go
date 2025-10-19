package encoders

import (
	"math"
	"time"
)

// RADecToAltAz converts right ascension and declination to altitude and azimuth
func RADecToAltAz(ra, dec, lat, lon float64, t time.Time) (alt, az float64) {
	// Convert RA (hours) to radians: 1 hour = 15 degrees
	raRad := ra * 15.0 * math.Pi / 180.0
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

	// Convert RA radians -> degrees -> hours (1 hour = 15 degrees)
	raDeg := ra * 180.0 / math.Pi
	raHours := raDeg / 15.0

	return raHours, dec
}

// calculateLST calculates Local Sidereal Time
func calculateLST(lon float64, t time.Time) float64 {
	// Calculate Local Sidereal Time using the algorithm from Jean Meeus
	// Steps:
	// 1) compute Julian Date (JD) for the UTC time
	// 2) compute number of days since J2000.0: D = JD - 2451545.0
	// 3) compute Greenwich Sidereal Time in degrees
	// 4) add longitude to obtain local sidereal time and convert to radians

	utc := t.UTC()

	year := utc.Year()
	month := int(utc.Month())
	day := utc.Day()
	hour := utc.Hour()
	min := utc.Minute()
	sec := utc.Second()
	nsec := utc.Nanosecond()

	// fractional day
	fracDay := (float64(hour) + float64(min)/60.0 + (float64(sec)+float64(nsec)/1e9)/3600.0) / 24.0

	y := year
	m := month
	d := float64(day) + fracDay

	if m <= 2 {
		y -= 1
		m += 12
	}

	A := int(math.Floor(float64(y) / 100.0))
	B := 2 - A + int(math.Floor(float64(A)/4.0))

	jd := math.Floor(365.25*float64(y+4716)) + math.Floor(30.6001*float64(m+1)) + d + float64(B) - 1524.5

	// days since J2000.0
	D := jd - 2451545.0
	T := D / 36525.0

	// Greenwich Sidereal Time in degrees
	gst := 280.46061837 + 360.98564736629*D + 0.000387933*T*T - (T*T*T)/38710000.0

	// normalize gst to [0,360)
	gst = math.Mod(gst, 360.0)
	if gst < 0 {
		gst += 360.0
	}

	// Local Sidereal Time in degrees
	lstDeg := gst + lon
	lstDeg = math.Mod(lstDeg, 360.0)
	if lstDeg < 0 {
		lstDeg += 360.0
	}

	// return in radians normalized to [0,2π)
	lstRad := lstDeg * math.Pi / 180.0
	lstRad = math.Mod(lstRad, 2*math.Pi)
	if lstRad < 0 {
		lstRad += 2 * math.Pi
	}
	return lstRad
}
