package test_cli

import (
	"testing"

	"github.com/cloudsftp/Sunangel/args"
	"github.com/cloudsftp/Sunangel/location"
)

func assertSunArgumentsEqual(t *testing.T, got, want *args.SunArguments) {
	if *got != *want && *got.Place != *want.Place {
		t.Errorf("got %v (loc: %v), want %v (loc: %v)", got, got.Place, want, want.Place)
	}
}

func TestParseCoordinates(t *testing.T) {
	arguments := []string{"cmd", "48.81", "9.58"}
	want := &args.SunArguments{
		Place:        location.NewLocation(48.81, 9.58),
		RadiusIgnore: 0,
		DayOffset:    0,
	}
	got, err := args.ParseSunArguments(arguments)
	if err != nil {
		panic(err)
	}

	assertSunArgumentsEqual(t, got, want)
}
