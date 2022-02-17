package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cloudsftp/Sunangel/sunset"
)

func TestSunsetEstimator(t *testing.T) {
	got := sunset.EstimateSunsetOf(dateCustom, locationGaensberg)
	want := time.Date(2022, time.February, 11, 17, 30, 47, 0, time.UTC)

	fmt.Println(got)

	assertDatePreciselyEqual(t, got, want)
}
