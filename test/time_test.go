package test

import (
	"testing"

	"github.com/cloudsftp/Sunangel/startime"
)

func TestHourAngleAt(t *testing.T) {
	got := startime.HourAngleAt(dateWiki, float64(11.6))
	want := float64(56.2387)

	assertApproxEqual(t, got, want)
}

func TestHourAngleAtCustom(t *testing.T) {
	got := startime.HourAngleAt(dateCustom, float64(9.58675))
	want := float64(38.806009069814)

	assertPreciselyEqual(t, got, want)
}

func TestJulianDaysSince2000At(t *testing.T) {
	got := startime.JulianDaysSince2000At(dateWiki)
	want := float64(2408.75)

	assertApproxEqual(t, got, want)
}

func TestJulianDaysSince2000AtCustom(t *testing.T) {
	got := startime.JulianDaysSince2000At(dateCustom)
	want := float64(8077.1875)

	assertApproxEqual(t, got, want)
}
