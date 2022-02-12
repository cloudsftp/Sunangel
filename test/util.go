package test

import (
	"math"
	"testing"
	"time"
)

var DateWiki = time.Date(2006, time.August, 6, 6, 0, 0, 0, time.UTC)
var DateCustom = time.Date(2022, time.February, 11, 17, 30, 0, 0, time.FixedZone("Berlin, DE", 3600))

func assertApproxEqualEpsilon(t *testing.T, got, want, epsilon float64) {
	difference := math.Abs(got - want)
	if difference > epsilon {
		t.Errorf("difference %.15f too big, got %.15f want %.15f", difference, got, want)
	}
}

func AssertApproxEqual(t *testing.T, got, want float64) {
	assertApproxEqualEpsilon(t, got, want, float64(0.5e-4))
}

func AssertPreciselyEqual(t *testing.T, got, want float64) {
	assertApproxEqualEpsilon(t, got, want, float64(0.5e-12))
}
