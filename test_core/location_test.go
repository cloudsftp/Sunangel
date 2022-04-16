package test_core

import (
	"math"
	"testing"

	"github.com/cloudsftp/Sunangel/location"
)

func testElevationGeneral(t *testing.T, loc *location.Location, want float64) {
	got := loc.GetElevation()

	assertPreciselyEqual(t, got, want)
}

func TestElevation(t *testing.T) {
	testElevationGeneral(t, locationParagleiter, 384)
	testElevationGeneral(t, locationTurbinesWTB, 478)
	testElevationGeneral(t, locationMuenchen, 540)
}

func testAngleToGeneral(t *testing.T, a, b *location.Location, want float64) {
	got := a.AltitudeAngleTo(b)

	assertPreciselyEqual(t, got, want)
}

func TestAngleTo(t *testing.T) {
	testAngleToGeneral(t, locationParagleiter, locationTurbinesWTB, 0.007488244041055)
	testAngleToGeneral(t, locationParagleiter, locationMuenchen, -0.012300042822072)
	testAngleToGeneral(t, locationMuenchen, locationParagleiter, -0.014151429022413)
}

func testAzimutAngleToGeneral(t *testing.T, a, b *location.Location, want float64) {
	got := a.AzimutAngleTo(b)

	assertApproxEqual(t, got, want)
}

func TestAzimutAngleTo(t *testing.T) {
	a := location.NewLocation(51.5, 0)
	b := location.NewLocation(-22.97, -43.18)

	want := float64(-2.4548) + 2*math.Pi

	testAzimutAngleToGeneral(t, a, b, want)
}
