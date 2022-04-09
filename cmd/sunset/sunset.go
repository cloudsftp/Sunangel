package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cloudsftp/Sunangel/args"
	"github.com/cloudsftp/Sunangel/horizon"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunset"
)

const dateLayout = "2006-01-02"
const timeLayout = "15:04:05 MST"

func main() {
	arguments, err := args.ParseSunArguments(os.Args)
	if err != nil {
		args.PrintSunUsage()
	}

	place := location.NewLocation(arguments.Place.Latitude, arguments.Place.Longitude)
	horizon := horizon.NewHorizon(place, arguments.StartRadius)

	date := time.Now().Add(time.Duration(arguments.DayOffset * 24 * int(time.Hour)))

	sunsetTime := sunset.EstimateSunsetOf(date, horizon)

	fmt.Printf(
		"\n\nResult:\nThe sun sets at %s on %s\n",
		sunsetTime.Format(timeLayout),
		sunsetTime.Format(dateLayout),
	)
}
