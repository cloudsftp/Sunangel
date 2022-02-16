package test

import (
	"testing"
	"time"

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

func testTimeOfDayAsDecimalGeneral(t *testing.T, date time.Time, want float64) {
	got := startime.TimeOfDayAsDecimal(date)

	assertPreciselyEqual(t, got, want)
}

func TestTimeOfDayAsDecimal(t *testing.T) {
	date := time.Date(0, time.January, 1, 6, 0, 0, 0, time.UTC)
	testTimeOfDayAsDecimalGeneral(t, date, float64(6))

	date = time.Date(0, time.January, 1, 0, 6, 0, 0, time.UTC)
	testTimeOfDayAsDecimalGeneral(t, date, float64(0.1))

	date = time.Date(0, time.January, 1, 0, 0, 12, 0, time.UTC)
	testTimeOfDayAsDecimalGeneral(t, date, float64(0.002))

	date = time.Date(0, time.January, 1, 0, 0, 0, 1000, time.UTC)
	testTimeOfDayAsDecimalGeneral(t, date, float64(0.0000000001))

	date = time.Date(0, time.January, 1, 16, 18, 24, 1000000, time.UTC)
	testTimeOfDayAsDecimalGeneral(t, date, float64(16.3040001))
}
