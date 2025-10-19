package encoders

import (
	"math"
	"testing"
	"time"
)

func TestRADecToAltAz(t *testing.T) {
	tests := []struct {
		ra, dec, lat, lon float64
		time              time.Time
		expectedAlt       float64
		expectedAz        float64
	}{
		{0, 0, 0, 0, time.Date(2024, 8, 24, 18, 0, 0, 0, time.UTC), -26.495355, 90},
		// RA: 90 degrees = 6 hours
		{6, 45, 41.89193, 12.51133, time.Date(2024, 8, 24, 18, 0, 0, 0, time.UTC), -2.213278, 350.153983},
		// Add more test cases here
		//Roma 41.89193 12.51133
	}

	for _, tt := range tests {
		alt, az := RADecToAltAz(tt.ra, tt.dec, tt.lat, tt.lon, tt.time)
		if math.Abs(alt-tt.expectedAlt) > 0.1 || math.Abs(az-tt.expectedAz) > 0.1 {
			t.Errorf("RADecToAltAz(%f, %f, %f, %f, %v) = (%f, %f), want (%f, %f)",
				tt.ra, tt.dec, tt.lat, tt.lon, tt.time, alt, az, tt.expectedAlt, tt.expectedAz)
		}
	}
}

func TestAltAzToRADec(t *testing.T) {
	tests := []struct {
		alt, az, lat, lon float64
		time              time.Time
		expectedRA        float64
		expectedDec       float64
	}{
		// expected RA given in hours (converted from degrees in original tests)
		{90, 0, 0, 0, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), 101.45 / 15.0, 0},
		{45, 270, 45, -75, time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC), 152.2 / 15.0, 30},
		// Add more test cases here
	}

	for _, tt := range tests {
		ra, dec := AltAzToRADec(tt.alt, tt.az, tt.lat, tt.lon, tt.time)
		if math.Abs(ra-tt.expectedRA) > 0.1 || math.Abs(dec-tt.expectedDec) > 0.1 {
			t.Errorf("AltAzToRADec(%f, %f, %f, %f, %v) = (%f, %f), want (%f, %f)",
				tt.alt, tt.az, tt.lat, tt.lon, tt.time, ra, dec, tt.expectedRA, tt.expectedDec)
		}
	}
}

func TestCalculateLST(t *testing.T) {
	tests := []struct {
		lon         float64
		time        time.Time
		expectedLST float64
	}{
		{0, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), 1.752159},
		{-75, time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC), 3.593357},
		// Add more test cases here
	}

	for _, tt := range tests {
		lst := calculateLST(tt.lon, tt.time)
		if math.Abs(lst-tt.expectedLST) > 0.0001 {
			t.Errorf("calculateLST(%f, %v) = %f, want %f", tt.lon, tt.time, lst, tt.expectedLST)
		}
	}
}
