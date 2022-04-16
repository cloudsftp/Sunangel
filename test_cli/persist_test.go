package test_cli

import (
	"testing"

	"github.com/cloudsftp/Sunangel/horizon"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/persist"
)

var testLocation = location.NewLocation(48.814, 9.58)

// Location

func TestAddAndGetLocations(t *testing.T) {
	name := "TestLocation"
	want := testLocation
	persist.AddLocation(name, want)

	locations := persist.GetLocations()
	got := locations["TestLocation"]

	if got == nil {
		t.Errorf("got nil")
		t.FailNow()
	}

	if *got != *want {
		t.Errorf("got %v, want %v", *got, *want)
	}
}

func TestAddAndGetLocation(t *testing.T) {
	name := "TestLocation2"
	want := testLocation
	persist.AddLocation(name, want)

	_, err := persist.GetLocation("TesstLocation")
	if err == nil {
		t.Error("expected error")
		t.FailNow()
	}

	got, err := persist.GetLocation("TestLocation2")
	if err != nil {
		panic(err)
	}

	if got == nil {
		t.Errorf("got nil")
		t.FailNow()
	}

	if *got != *want {
		t.Errorf("got %v, want %v", *got, *want)
	}
}

func TestAddAndDeleteLocation(t *testing.T) {
	name := "TestLocation3"
	persist.AddLocation(name, testLocation)

	persist.DeleteLocation(name)

	_, err := persist.GetLocation(name)
	if err == nil {
		t.Error("expected error")
	}
}

// Horizon

func assertAltitudeArrayEqual(t *testing.T, got, want horizon.AltitudeArray) {
	if len(got) != len(want) {
		t.Errorf("altitude array length mismatch got %d want %d", len(got), len(want))
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("altitude array element %d mismatch got %f want %f", i, got[i], want[i])
		}
	}
}

func assertHorizonEqual(t *testing.T, got, want *horizon.Horizon) {
	if got.Place.Latitude != got.Place.Latitude ||
		got.Place.Longitude != want.Place.Longitude ||
		got.GetStartRadius() != want.GetStartRadius() {
		t.Errorf("got %v (%v), want %v (%v)", *got, *got.Place, *want, *want.Place)
		t.FailNow()
	}

	assertAltitudeArrayEqual(t, got.GetAltitudeArray(), want.GetAltitudeArray())
}

func TestAddAndGetHorizon(t *testing.T) {
	alt := horizon.AltitudeArray{0, 1, 2}
	radius := 2
	want := horizon.NewHorizonWithAltitude(testLocation, radius, alt)

	persist.AddHorizon(want)

	got, err := persist.GetHorizon(testLocation, radius)
	if err != nil {
		panic(err)
	}

	assertHorizonEqual(t, got, want)
}

func TestDeleteHorizonAll(t *testing.T) {
	startRadius := 3
	hor := horizon.NewHorizonWithAltitude(testLocation, startRadius, horizon.AltitudeArray{})
	persist.AddHorizon(hor)

	persist.DeleteHorizonAll()

	_, err := persist.GetHorizon(testLocation, startRadius)
	if err == nil {
		t.Error("expected error")
	}
}
