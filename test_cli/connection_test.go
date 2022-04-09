package test_cli

import (
	"testing"
	"time"

	"github.com/cloudsftp/Sunangel/horizon"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunset"
)

func TestLocation(t *testing.T) {
	var loc *location.Location = location.NewLocation(48.814, 9.11)

	var hor *horizon.Horizon = horizon.NewHorizon(loc, 500)

	date := time.Now()
	date = sunset.EstimateSunsetOf(date, hor)

	_ = date
}
