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
		t.Errorf("got %04d-%02d-%02d %02d:%02d:%02d want %04d-%02d-%02d %02d:%02d:%02d",
			got.Year(), got.Month(), got.Day(), got.Hour(), got.Minute(), got.Second(),
			want.Year(), want.Month(), want.Day(), want.Hour(), want.Minute(), want.Second(),
		)
	}
}

func testSunsetEstimatorGeneral(t *testing.T, wantString string, place *location.Location) {
	want, err := time.Parse(timeLayout, wantString)
	if err != nil {
		t.Errorf("could not parse date '%s'", wantString)
	}

	horizon := horizon.NewHorizon(place)
	got := sunset.EstimateSunsetOf(want, horizon)

	assertDatePreciselyEqual(t, got, want)
}

func TestSunsetEstimatorParagleiter(t *testing.T) {
	testSunsetEstimatorGeneral(t, "2022-03-23 18:34:00 CET", locationParagleiter)
}

func TestSunsetEstimatorFreibad(t *testing.T) {
	testSunsetEstimatorGeneral(t, "2022-03-22 18:25:29 CET", locationFreibad)
}

func TestSunsetEstimatorOWH(t *testing.T) {
	testSunsetEstimatorGeneral(t, "2022-03-24 18:34:50 CET", locationOWH)
	testSunsetEstimatorGeneral(t, "2022-04-03 19:47:23 CEST", locationOWH)
}

func TestSunsetEstimatorPluederwiese(t *testing.T) { // TODO: folter out near hills
	testSunsetEstimatorGeneral(t, "2022-03-26 18:43:38 CET", locationPluederwiese)
}
