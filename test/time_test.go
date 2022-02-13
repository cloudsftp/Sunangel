package test

import (
	"math"
	"testing"

	"github.com/cloudsftp/Sunangel/startime"
)

func TestHourAngleOfSunAt(t *testing.T) {
	got := startime.HourAngleOfSunAt(dateWiki, longitudeWiki)
	want := float64((2 * math.Pi) - 1.394169)

	assertApproxEqual(t, got, want)
}
