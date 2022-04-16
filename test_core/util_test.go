package test_core

import (
	"testing"

	"github.com/cloudsftp/Sunangel/util"
)

func TestLinInt(t *testing.T) {
	got := util.LinInt(6, 5, 10, 9, 18)
	want := float64(12)

	assertPreciselyEqual(t, got, want)

	got = util.LinInt(0.2, 0, 0.5, 0.6, 2)
	want = float64(1)

	assertPreciselyEqual(t, got, want)
}
