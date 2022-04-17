package test_core

import (
	"math"
	"testing"
	"time"

	"github.com/cloudsftp/Sunangel/location"
)

var (
	dateWiki         = time.Date(2006, time.August, 6, 6, 0, 0, 0, time.UTC)
	locationMuenchen = location.NewLocation(48.1, 11.6)

	dateCustom          = time.Date(2022, time.February, 11, 17, 30, 0, 0, time.FixedZone("Berlin, DE", 3600))
	locationParagleiter = location.NewLocation(48.8187132, 9.5878127)

	locationTurbinesWTB = location.NewLocation(48.7866067, 9.4424222)
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
