package test

import (
	"testing"
	"time"

	"github.com/cloudsftp/Sunangel/sunset"
)

func TestSunsetEstimator(t *testing.T) {
	got := sunset.EstimateSunsetOf(dateCustom, locationGaensberg)
	want := time.Date(2022, time.February, 11, 17, 30, 41, 0, time.UTC)

	assertDatePreciselyEqual(t, got, want)
}
