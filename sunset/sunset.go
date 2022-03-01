package sunset

import (
	"log"
	"time"

	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunangel"
)

// EstimateSunsetOf returns an estimate of the time of sunset
// at a given time and place.
func EstimateSunsetOf(date time.Time, place *location.Location) time.Time {
	year := date.Year()
	month := date.Month()
	day := date.Day()
	loc := date.Location()

	lowerBound := time.Date(year, month, day, 12, 0, 0, 0, loc)
	upperBound := time.Date(year, month, day, 23, 59, 59, 1e9-1, loc)

	return binarySunsetSearch(lowerBound, upperBound, place)
}

func binarySunsetSearch(lowerBound, upperBound time.Time, place *location.Location) time.Time {
	place.LoadHorizonOfLocation()

	limitSearchDuration := time.Duration(1e9)

	for {
		currentSearchDuration := upperBound.Sub(lowerBound)
		newBound := lowerBound.Add(currentSearchDuration / 2)

		log.Print(newBound)

		azimutAngle := sunangel.AzimutSunAngleAt(newBound, place)
		horizonAngle := sunangel.AltitudeSunAngleAt(newBound, place)

		horizonAngleGoal := place.GetHorizonAngleAt(azimutAngle)

		log.Printf("got %f want %f", horizonAngle, horizonAngleGoal)

		if horizonAngle < horizonAngleGoal {
			upperBound = newBound
		} else {
			lowerBound = newBound
		}

		if currentSearchDuration < limitSearchDuration {
			break
		}
	}

	return lowerBound.Round(limitSearchDuration)
}
