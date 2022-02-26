package test

import (
	"testing"

	"github.com/cloudsftp/Sunangel/location"
)

func testElevationGeneral(t *testing.T, loc location.Location, want float64) {
	got := loc.GetElevation()

	assertPreciselyEqual(t, got, want)
}

func TestElevation(t *testing.T) {
	testElevationGeneral(t, locationGaensberg, 384)
	testElevationGeneral(t, locationTurbinesWTB, 478)
	testElevationGeneral(t, locationMuenchen, 540)
}

func testAngleToGeneral(t *testing.T, a, b location.Location, want float64) {
	got := a.HorizontalAngleTo(b)

	assertPreciselyEqual(t, got, want)
}

func TestAngleTo(t *testing.T) {
	testAngleToGeneral(t, locationGaensberg, locationTurbinesWTB, 0.004382285396)
	testAngleToGeneral(t, locationGaensberg, locationMuenchen, -0.017952603423)
	testAngleToGeneral(t, locationMuenchen, locationGaensberg, -0.019268073648)
}
