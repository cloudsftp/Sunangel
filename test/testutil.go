package test

import (
	"testing"

	"github.com/cloudsftp/Sunangel/util"
)

var (
	dateWiki   = util.DateWiki
	dateCustom = util.DateCustom
)

func assertApproxEqual(t *testing.T, got, want float64) {
	util.AssertApproxEqual(t, got, want)
}

func assertPreciselyEqual(t *testing.T, got, want float64) {
	util.AssertPreciselyEqual(t, got, want)
}
