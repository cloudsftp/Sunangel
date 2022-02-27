package test

import (
	"testing"
	"time"

	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunset"
)

const timeLayout string = "2006-01-02 15:04:05"

var berlinTiomezone = time.FixedZone("Berlin, DE", 3600)

func assertDatePreciselyEqual(t *testing.T, got, want time.Time) {
	if got.Year() != want.Year() || got.Month() != want.Month() || got.Day() != want.Day() ||
		got.Hour() != want.Hour() || got.Minute() != want.Minute() || got.Second() != want.Second() {
		t.Errorf("got %04d-%02d-%02d %02d:%02d:%02d want %04d-%02d-%02d %02d:%02d:%02d",
			got.Year(), got.Month(), got.Day(), got.Hour(), got.Minute(), got.Second(),
			want.Year(), want.Month(), want.Day(), want.Hour(), want.Minute(), want.Second(),
		)
	}
}

func testSunsetEstimatorGeneral(t *testing.T, wantString string, place location.Location) {
	want, err := time.ParseInLocation(timeLayout, wantString, berlinTiomezone)
	if err != nil {
		t.Errorf("could not parse date '%s'", wantString)
	}

	got := sunset.EstimateSunsetOf(want, place)

	assertDatePreciselyEqual(t, got, want)
}

// These tests depend on the horizon
func TestSunsetEstimatorFreibad(t *testing.T) {
	place := *location.NewLocation(48.8292463, 9.5773359)
	place.RecomputeHorizon()

	testSunsetEstimatorGeneral(t, "2022-02-27 17:55:00", place)
}

// These tests don't depend on the horizon

func TestSunsetEstimatorGaensberg(t *testing.T) {
	testSunsetEstimatorGeneral(t, "2022-02-11 17:30:41", locationGaensberg)
}
