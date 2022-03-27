package vis

import (
	"fmt"
	"time"

	"github.com/arafatk/glot"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/util"

	"github.com/cloudsftp/Sunangel/sunangel"
	"github.com/cloudsftp/Sunangel/sunset"
)

func VisualizeSunset(place *location.Location, date time.Time) {
	sunsetTime := sunset.EstimateSunsetOf(date, place)

	duration := 6 * time.Hour
	startTime := sunsetTime.Add(-duration / 2)
	VisualizeHorizon(place, startTime, duration)
}

func VisualizeHorizon(loc *location.Location, startTime time.Time, duration time.Duration) {
	plot, err := glot.NewPlot(2, false, false)
	if err != nil {
		panic(err)
	}

	timeStep := time.Second
	numSteps := int(duration / timeStep)

	azimut := make([]float64, numSteps)
	altitude := make([]float64, numSteps)

	currTime := startTime
	for i := 0; i < numSteps; i++ {
		azimut[i] = sunangel.AzimutSunAngleAt(currTime, loc)
		altitude[i] = sunangel.AltitudeSunAngleAt(currTime, loc)

		currTime = currTime.Add(timeStep)
	}

	x := make([]float64, numSteps)

	startAngle := azimut[0]
	stopAngle := azimut[len(azimut)-1]
	angleStep := (stopAngle - startAngle) / float64(numSteps)
	currAngle := startAngle
	for i := 0; i < numSteps; i++ {
		x[i] = currAngle
		currAngle += angleStep
	}

	altitudeAtAzimut := func(azimutVal float64) float64 {
		leftIndex := 0
		for leftIndex < len(azimut)-2 && azimut[leftIndex] < azimutVal {
			leftIndex++
		}

		return util.LinInt(
			azimutVal,
			azimut[leftIndex], altitude[leftIndex],
			azimut[leftIndex+1], altitude[leftIndex+1],
		)
	}
	plot.AddFunc2d("Sun", "lines", x, altitudeAtAzimut)

	horizonAltitudeAtAzimut := func(azimutVal float64) float64 {
		return loc.GetHorizonAngleAt(azimutVal)
	}
	plot.AddFunc2d("Horizon", "lines", x, horizonAltitudeAtAzimut)

	plot.SavePlot(fmt.Sprintf("Img/Horizons/%s.png", loc.Name))

}
