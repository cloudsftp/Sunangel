package test

import (
	"math"
	"testing"

	"github.com/cloudsftp/Sunangel/angle"
	"github.com/cloudsftp/Sunangel/sunangel"
)

func TestAltitudeSunangleAt(t *testing.T) {
	got := sunangel.AltitudeSunAngleAt(dateWiki, locationMuenchen)
	want := angle.RadiansFromDegrees(19.110)

	assertApproxEqual(t, got, want)
}

func TestAzimutSunAngleAt(t *testing.T) {
	got := sunangel.AzimutSunAngleAt(dateWiki, locationMuenchen)
	want := angle.RadiansFromDegrees(265.938)
	want += math.Pi
	want = angle.NormalizeRadiansLatitude(want)

	assertApproxEqual(t, got, want)
}

func TestAltitudeSunangleAtCustom(t *testing.T) {
	got := sunangel.AltitudeSunAngleAt(dateCustom, locationGaensberg)
	want := float64(0.00902)

	assertApproxEqual(t, got, want)
}

func TestAzimutSunAngleAtCustom(t *testing.T) {
	got := sunangel.AzimutSunAngleAt(dateCustom, locationGaensberg)
	want := float64(1.19716) + math.Pi

	assertApproxEqual(t, got, want)
}
