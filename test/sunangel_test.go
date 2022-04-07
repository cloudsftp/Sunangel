package test

import (
	"math"
	"testing"
	"time"

	"github.com/cloudsftp/Sunangel/angle"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunangel"
)

var berlinTiomezone = time.FixedZone("Berlin, DE", 3600)

func testAltitudeAngleGeneral(t *testing.T, date time.Time, place *location.Location, want float64) {
	got := sunangel.AltitudeSunAngleAt(date, place)

	assertApproxEqual(t, got, want)
}

func TestAltitudeSunangleAt(t *testing.T) {
	want := angle.RadiansFromDegrees(19.110)
	testAltitudeAngleGeneral(t, dateWiki, locationMuenchen, want)
}

func TestAzimutSunAngleAt(t *testing.T) {
	got := sunangel.AzimutSunAngleAt(dateWiki, locationMuenchen)
	want := angle.RadiansFromDegrees(265.938)
	want += math.Pi
	want = angle.NormalizeRadiansLatitude(want)

	assertApproxEqual(t, got, want)
}

func TestAltitudeSunangleAtCustom(t *testing.T) {
	testAltitudeAngleGeneral(t, dateCustom, locationParagleiter, 0.00902)
}

func TestAzimutSunAngleAtCustom(t *testing.T) {
	got := sunangel.AzimutSunAngleAt(dateCustom, locationParagleiter)
	want := float64(1.19716) + math.Pi

	assertApproxEqual(t, got, want)
}

func TestAltitudeAngleAtBugNegativeGaensberg(t *testing.T) {
	date := time.Date(2022, time.February, 27, 17, 57, 00, 703124999999, berlinTiomezone)
	want := float64(-0.02128)

	testAltitudeAngleGeneral(t, date, locationParagleiter, want)
}
