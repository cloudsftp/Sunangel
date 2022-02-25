package main

import (
	"fmt"
	"time"

	"github.com/cloudsftp/Sunangel/angle"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunangel"
	"github.com/cloudsftp/Sunangel/sunset"
)

func main() {
	fmt.Println("Sunset today at Gaensberg w/o horizon calc")
	date := time.Now()
	locationGaensberg := *location.NewLocation(48.8187132, 9.5878127)
	estimatedSunsetTime := sunset.EstimateSunsetOf(date, locationGaensberg)
	fmt.Println(estimatedSunsetTime)

	fmt.Println("\nMISC")
	fmt.Println(angle.DegreesFromRadians(sunangel.AzimutSunAngleAt(date, locationGaensberg)))

	locationGaensberg.ComputeHorizon()
}
