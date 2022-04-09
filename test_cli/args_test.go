package test_cli

import (
	"testing"

	"github.com/cloudsftp/Sunangel/args"
)

func assertSunArgumentsEqual(t *testing.T, got, want *args.SunArguments) {
	if got.Place.Latitude != want.Place.Latitude ||
		got.Place.Longitude != want.Place.Longitude ||
		got.StartRadius != want.StartRadius ||
		got.DayOffset != want.DayOffset {
		t.Errorf("got %v (loc: %v), want %v (loc: %v)", got, got.Place, want, want.Place)
	}
}

func testParseSunArgumentsGeneral(t *testing.T, arguments []string, latitude, longitude float64, startRadius, dayOffset int) {
	want := args.NewSunArguments(latitude, longitude, startRadius, dayOffset)
	got, err := args.ParseSunArguments(arguments)
	if err != nil {
		panic(err)
	}

	assertSunArgumentsEqual(t, got, want)
}

func TestParseCoordinates(t *testing.T) {
	arguments := []string{"cmd", "48.81", "9.58"}
	testParseSunArgumentsGeneral(t, arguments, 48.81, 9.58, 0, 0)
}

func TestParseCoordinatesDayOffset(t *testing.T) {
	arguments := []string{"cmd", "48.81", "9.58", "d=20"}
	testParseSunArgumentsGeneral(t, arguments, 48.81, 9.58, 0, 20)
}

func TestParseCoordinatesStartRadius(t *testing.T) {
	arguments := []string{"cmd", "48.81", "9.58", "r=20"}
	testParseSunArgumentsGeneral(t, arguments, 48.81, 9.58, 20, 0)
}

func TestParseCoordinatesStartRadiusAndDayOffset(t *testing.T) {
	arguments := []string{"cmd", "48.81", "9.58", "r=20", "d=30"}
	testParseSunArgumentsGeneral(t, arguments, 48.81, 9.58, 20, 30)

	arguments = []string{"cmd", "48.81", "9.58", "d=20", "r=30"}
	testParseSunArgumentsGeneral(t, arguments, 48.81, 9.58, 30, 20)
}
