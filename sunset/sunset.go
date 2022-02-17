package sunset

import (
	"time"

	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/sunangel"
)

// EstimateSunsetOf returns an estimate of the time of sunset
// at a given time and place.
func EstimateSunsetOf(date time.Time, place location.Location) time.Time {
	year := date.Year()
	month := date.Month()
	day := date.Day()
	loc := date.Location()

	lowerBound := time.Date(year, month, day, 12, 0, 0, 0, loc)
	upperBound := time.Date(year, month, day, 23, 59, 59, 1e9-1, loc)

	return simpleBinarySunsetSearch(lowerBound, upperBound, place)
}

func simpleBinarySunsetSearch(lowerBound, upperBound time.Time, place location.Location) time.Time {
	done := false
	sunAngleGoal := float64(0)

	limitSearchDuration := time.Duration(1e9)

	var newBound time.Time
	var currentSearchDuration time.Duration
	var sunAngleNewBound float64
	for !done {
		currentSearchDuration = upperBound.Sub(lowerBound)
		if currentSearchDuration/2 < limitSearchDuration {
			done = true
		}

		newBound = lowerBound.Add(currentSearchDuration / 2)

		sunAngleNewBound = sunangel.AltitudeSunAngleAt(newBound, place)
		if sunAngleNewBound < sunAngleGoal {
			upperBound = newBound
		} else {
			lowerBound = newBound
		}
	}

	result := lowerBound
	result = result.Round(limitSearchDuration)
	return result
}
