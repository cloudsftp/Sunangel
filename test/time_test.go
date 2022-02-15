package test

import (
	"testing"

	"github.com/cloudsftp/Sunangel/startime"
)

func TestJulianDaysSince2000At(t *testing.T) {
	got := startime.JulianDaysSince2000At(dateWiki)
	want := float64(2408.75)

	assertPreciselyEqual(t, got, want)
}

func TestJulianDaysSince2000ToMidnightOf(t *testing.T) {
	got := startime.JulianDaysSince2000ToMidnightOf(dateWiki)
	want := float64(2408.5)

	assertPreciselyEqual(t, got, want)
}

func TestJulianCenturiesSince2000ToMidnightOf(t *testing.T) {
	got := startime.JulianCenturiesSince2000ToMidnightOf(dateWiki)
	want := float64(0.06594113621)

	assertApproxEqual(t, got, want)
}
