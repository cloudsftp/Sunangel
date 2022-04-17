package test_cli

import (
	"testing"

	"github.com/cloudsftp/Sunangel/cmd/args"
	"github.com/cloudsftp/Sunangel/location"
)

const (
	name    = "locname"
	lat     = 48.81
	latstr  = "48.81"
	long    = 9.58
	longstr = "9.58"
)

// Sun Arguments

func assertSunArgumentsEqual(t *testing.T, got, want *args.SunArguments) {
	ok := got.StartRadius == want.StartRadius &&
		got.DayOffset == want.DayOffset

	if got.Mode != want.Mode {
		ok = false
	} else {
		switch got.Mode {
		case args.Coordinates:
			ok = ok &&
				got.Place.Latitude == want.Place.Latitude &&
				got.Place.Longitude == want.Place.Longitude
		case args.Name:
			ok = ok && got.Name == want.Name
		default:
			ok = false
		}
	}

	if !ok {
		t.Errorf("got %v (loc: %v), want %v (loc: %v)", got, got.Place, want, want.Place)
	}
}

func testParseSunArgumentsGeneral(t *testing.T, arguments []string, want *args.SunArguments) {
	got, err := args.ParseSunArguments(arguments)
	if err != nil {
		t.Errorf("an error occured while parsing sun arguments: %v", err)
		t.FailNow()
	}

	assertSunArgumentsEqual(t, got, want)
}

func testParseSunArgumentsError(t *testing.T, arguments []string) {
	_, err := args.ParseSunArguments(arguments)
	if err == nil {
		t.Error("expected error")
	}
}

func testParseSunCoordinatesGeneral(t *testing.T, arguments []string, latitude, longitude float64, startRadius, dayOffset int) {
	want := &args.SunArguments{
		Mode:        args.Coordinates,
		Place:       location.NewLocation(latitude, longitude),
		StartRadius: startRadius,
		DayOffset:   dayOffset,
	}

	testParseSunArgumentsGeneral(t, arguments, want)
}

func TestParseCoordinates(t *testing.T) {
	arguments := []string{"cmd", "48.81", "9.58"}
	testParseSunCoordinatesGeneral(t, arguments, 48.81, 9.58, 0, 0)
}

func TestParseCoordinatesDayOffset(t *testing.T) {
	arguments := []string{"cmd", "48.81", "9.58", "d=20"}
	testParseSunCoordinatesGeneral(t, arguments, 48.81, 9.58, 0, 20)
}

func TestParseCoordinatesStartRadius(t *testing.T) {
	arguments := []string{"cmd", "48.81", "9.58", "r=20"}
	testParseSunCoordinatesGeneral(t, arguments, 48.81, 9.58, 20, 0)
}

func TestParseCoordinatesStartRadiusAndDayOffset(t *testing.T) {
	arguments := []string{"cmd", "48.81", "9.58", "r=20", "d=30"}
	testParseSunCoordinatesGeneral(t, arguments, 48.81, 9.58, 20, 30)

	arguments = []string{"cmd", "48.81", "9.58", "d=20", "r=30"}
	testParseSunCoordinatesGeneral(t, arguments, 48.81, 9.58, 30, 20)
}

func TestParseSunCoordinatesError(t *testing.T) {
	testParseSunArgumentsError(t, []string{"cmd", "48,81", "9.58"})
	testParseSunArgumentsError(t, []string{"cmd", "48.81", "9,58"})
	testParseSunArgumentsError(t, []string{"cmd", "48.81"})
}

func testParseSunNameGeneral(t *testing.T, arguments []string, name string, startRadius, dayOffset int) {
	want := &args.SunArguments{
		Mode:        args.Name,
		Name:        name,
		StartRadius: startRadius,
		DayOffset:   dayOffset,
	}

	testParseSunArgumentsGeneral(t, arguments, want)
}

func TestParseSunName(t *testing.T) {
	name := "locationname"
	arguments := []string{"cmd", name}
	testParseSunNameGeneral(t, arguments, name, 0, 0)
}

func TestParseSunNameStartRadius(t *testing.T) {
	name := "locationname"
	arguments := []string{"cmd", name, "r=30"}
	testParseSunNameGeneral(t, arguments, name, 30, 0)
}

func TestParseSunNameDayOffset(t *testing.T) {
	name := "locationname"
	arguments := []string{"cmd", name, "d=30"}
	testParseSunNameGeneral(t, arguments, name, 0, 30)
}

// Location Arguments

func assertLocationArgumentsEqual(t *testing.T, got, want *args.LocationArguments) {
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func testLocationArgumentsGeneral(t *testing.T, arguments []string, want *args.LocationArguments) {
	got, err := args.ParseLocationArguments(arguments)
	if err != nil {
		t.Errorf("error occured while parsing location arguments: %v", err)
		t.FailNow()
	}

	assertLocationArgumentsEqual(t, got, want)
}

func testLocationArgumentsError(t *testing.T, arguments []string) {
	_, err := args.ParseLocationArguments(arguments)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestParseLocationAdd(t *testing.T) {
	arguments := []string{"cmd", "add", name, latstr, longstr}
	want := &args.LocationArguments{
		Mode:      args.Add,
		Name:      name,
		Latitude:  lat,
		Longitude: long,
	}
	testLocationArgumentsGeneral(t, arguments, want)

	testLocationArgumentsError(t, []string{"cmd", "add", name, latstr})
	testLocationArgumentsError(t, []string{"cmd", "add", name})
	testLocationArgumentsError(t, []string{"cmd", "add"})
	testLocationArgumentsError(t, []string{"cmd", "add", name, latstr, "burr"})
}

func TestParseLocationList(t *testing.T) {
	arguments := []string{"cmd", "list"}
	want := &args.LocationArguments{Mode: args.List}
	testLocationArgumentsGeneral(t, arguments, want)

	testLocationArgumentsError(t, []string{"cmd", "list", "bla"})
}

func TestParseLocationDelete(t *testing.T) {
	arguments := []string{"cmd", "delete", name}
	want := &args.LocationArguments{
		Mode: args.Delete,
		Name: name,
	}
	testLocationArgumentsGeneral(t, arguments, want)

	testLocationArgumentsError(t, []string{"cmd", "delete"})
	testLocationArgumentsError(t, []string{"cmd", "delete", "a", "b"})
}

func TestParseLocationWrongCommand(t *testing.T) {
	testLocationArgumentsError(t, []string{"cmd", name})
	testLocationArgumentsError(t, []string{"cmd"})
}
