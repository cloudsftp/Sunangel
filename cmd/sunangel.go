package main

import (
	"fmt"
	"time"

	"github.com/cloudsftp/Sunangel/angle"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunangel"
	"github.com/cloudsftp/Sunangel/sunset"
	"github.com/cloudsftp/Sunangel/vis"
)

func main() {

	date := time.Now().Add(-0 * 24 * time.Hour)
	// loc := location.NewLocation("Paragleiter", 48.8187132, 9.5878127)
	// loc := location.NewLocation("Burg", 48.8230357, 9.5823731)
	// loc := location.NewLocation("Freibad", 48.8292463, 9.5773359)
	// loc := location.NewLocation("Freibad oben", 48.83523, 9.57165)
	loc := location.NewLocation("OWH", 48.814, 9.59172)
	// loc := location.NewLocation("Pluderwiese", 48.8320969, 9.6042998)

	// loc.RecomputeHorizon()

	estimatedSunsetTime := sunset.EstimateSunsetOf(date, loc)
	azimutAngle := sunangel.AzimutSunAngleAt(estimatedSunsetTime, loc)
	horizonAngle := loc.GetHorizonAngleAt(azimutAngle)

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

	vis.VisualizeSunset(loc, date)

}
