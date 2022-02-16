package test

import (
	"math"
	"testing"
	"time"
)

var (
	dateWiki              = time.Date(2006, time.August, 6, 6, 0, 0, 0, time.UTC)
	latitudeWiki  float64 = 0.839503
	longitudeWiki float64 = 0.202458

	dateCustom              = time.Date(2022, time.February, 11, 17, 30, 0, 0, time.FixedZone("Berlin, DE", 3600))
	latitudeCustom  float64 = 0.852043560884
	longitudeCustom float64 = 0.167320701131
)

func assertApproxEqualEpsilon(t *testing.T, got, want, epsilon float64) {
	difference := math.Abs(got - want)
	if difference > epsilon {
		t.Errorf("difference %.15f too big, got %.15f want %.15f", difference, got, want)
	}
}

func assertApproxEqual(t *testing.T, got, want float64) {
	assertApproxEqualEpsilon(t, got, want, float64(0.5e-4))
}

func assertPreciselyEqual(t *testing.T, got, want float64) {
	assertApproxEqualEpsilon(t, got, want, float64(0.5e-12))
}

func assertEqual(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.15f want %.15f", got, want)
	}
}
