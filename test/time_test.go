package test

import (
	"testing"

	"github.com/cloudsftp/Sunangel/time"
)

func TestHourAngleAt(t *testing.T) {
	got := time.HourAngleAt(DateWiki, float64(11.6))
	want := float64(56.2387)

	AssertApproxEqual(t, got, want)
}

func TestHourAngleAtCustom(t *testing.T) {
	got := time.HourAngleAt(DateCustom, float64(9.58675))
	want := float64(38.806009069814)

	AssertPreciselyEqual(t, got, want)
}
