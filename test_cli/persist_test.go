package test_cli

import (
	"testing"

	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/persist"
)

var testLocation = location.NewLocation(48.814, 9.58)

func TestStoreAndLoadLocations(t *testing.T) {
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

func TestStoreAndLoadLocation(t *testing.T) {
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

func TestStoreAndDeleteLocation(t *testing.T) {
	name := "TestLocation3"
	persist.AddLocation(name, testLocation)

	persist.DeleteLocation(name)

	_, err := persist.GetLocation(name)
	if err == nil {
		t.Error("expected error")
	}
}
