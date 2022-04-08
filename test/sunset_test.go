package test

import (
	"testing"
	"time"

	"github.com/cloudsftp/Sunangel/horizon"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunset"
)

const timeLayout string = "2006-01-02 15:04:05 MST"

var (
	locationFreibad      = location.NewLocation("Freibad", 48.8292463, 9.5773359)
	locationOWH          = location.NewLocation("OWH", 48.814, 9.59172)
	locationPluederwiese = location.NewLocation("Plüderwiese", 48.8320969, 9.6042998)
)

func assertDatePreciselyEqual(t *testing.T, got, want time.Time) {
	if got.Year() != want.Year() || got.Month() != want.Month() || got.Day() != want.Day() ||
		got.Hour() != want.Hour() || got.Minute() != want.Minute() || got.Second() != want.Second() {
		t.Errorf("got %s want %s", got.Format(timeLayout), want.Format(timeLayout))
	}
}

func testSunsetEstimatorGeneral(t *testing.T, wantString string, horizon *horizon.Horizon) {
	want, err := time.Parse(timeLayout, wantString)
	if err != nil {
		t.Errorf("could not parse date '%s'", wantString)
	}

	got := sunset.EstimateSunsetOf(want, horizon)

	assertDatePreciselyEqual(t, got, want)
}

func TestSunsetEstimatorParagleiter(t *testing.T) {
	horizonParagleiter := horizon.NewHorizon(locationParagleiter, 0)
	testSunsetEstimatorGeneral(t, "2022-03-23 18:33:58 CET", horizonParagleiter)
}

func TestSunsetEstimatorFreibad(t *testing.T) {
	horizonFreibad := horizon.NewHorizon(locationFreibad, 0)
	testSunsetEstimatorGeneral(t, "2022-03-22 18:25:30 CET", horizonFreibad)
}

func TestSunsetEstimatorOWH(t *testing.T) {
	horizonOWH := horizon.NewHorizon(locationOWH, 0)
	testSunsetEstimatorGeneral(t, "2022-03-24 18:34:43 CET", horizonOWH)
	testSunsetEstimatorGeneral(t, "2022-04-03 19:47:26 CEST", horizonOWH)
}

func TestSunsetEstimatorPluederwiese(t *testing.T) {
	horizonPluederwiese := horizon.NewHorizon(locationPluederwiese, 500)
	testSunsetEstimatorGeneral(t, "2022-03-26 18:33:36 CET", horizonPluederwiese)
}
