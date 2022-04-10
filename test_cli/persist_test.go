package test_cli

import (
	"testing"

	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/persist"
)

func TestStoreAndLoadLocations(t *testing.T) {
	name := "TestLocation"
	want := location.NewLocation(48.814, 9.58)
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
	want := location.NewLocation(48.815, 9.57)
	persist.AddLocation(name, want)

	_, err := persist.GetLocation("TesstLocation")
	if err == nil {
		t.Error("expected error")
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
