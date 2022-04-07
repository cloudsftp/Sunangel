package main

import (
	"fmt"
	"time"

	"github.com/cloudsftp/Sunangel/angle"
	"github.com/cloudsftp/Sunangel/horizon"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunangel"
	"github.com/cloudsftp/Sunangel/sunset"
	"github.com/cloudsftp/Sunangel/vis"
)

func main() {

	date := time.Now().Add(-0 * 24 * time.Hour)
	// place := location.NewLocation("Paragleiter", 48.8187132, 9.5878127)
	// place := location.NewLocation("Burg", 48.8230357, 9.5823731)
	// place := location.NewLocation("Freibad", 48.8292463, 9.5773359)
	// place := location.NewLocation("Freibad oben", 48.83523, 9.57165)
	place := location.NewLocation("OWH", 48.814, 9.59172)
	// place := location.NewLocation("Pluderwiese", 48.8320969, 9.6042998)

	horizon := horizon.NewHorizon(place)

	estimatedSunsetTime := sunset.EstimateSunsetOf(date, horizon)
	azimutAngle := sunangel.AzimutSunAngleAt(estimatedSunsetTime, place)
	horizonAngle := horizon.GetAltitude(azimutAngle)

	fmt.Println(estimatedSunsetTime)
	fmt.Printf(
		"azimut %f altitute %f\n",
		angle.DegreesFromRadians(azimutAngle),
		angle.DegreesFromRadians(horizonAngle),
	)
	/*
		fmt.Print(loc.Horizon)
		for i := 0; i < len(loc.Horizon); i++ {
			fmt.Printf("index %4d, angle %f\n", i, angle.DegreesFromRadians(loc.Horizon[i]))
		}*/

	vis.VisualizeSunset(horizon, date)

}
