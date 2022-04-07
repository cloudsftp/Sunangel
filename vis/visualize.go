package vis

import (
	"fmt"
	"math"
	"time"

	"github.com/arafatk/glot"
	"github.com/cloudsftp/Sunangel/angle"
	"github.com/cloudsftp/Sunangel/horizon"

	"github.com/cloudsftp/Sunangel/sunangel"
	"github.com/cloudsftp/Sunangel/sunset"
)

func VisualizeSunset(horizon *horizon.Horizon, date time.Time) {
	sunsetTime := sunset.EstimateSunsetOf(date, horizon)

	duration := 6 * time.Hour
	startTime := sunsetTime.Add(-duration / 2)
	VisualizeHorizon(horizon, startTime, duration)
}

func VisualizeHorizon(horizon *horizon.Horizon, startTime time.Time, duration time.Duration) {
	plot, err := glot.NewPlot(2, false, false)
	if err != nil {
		panic(err)
	}

	timeStep := time.Second
	numSteps := int(duration / timeStep)

	azimut := make([]float64, numSteps)
	altitude := make([]float64, numSteps)

	lastAzimut := float64(0)
	azimutOffset := float64(0)

	currTime := startTime
	for i := 0; i < numSteps; i++ {
		currAzimut := sunangel.AzimutSunAngleAt(currTime, horizon.Place)
		if currAzimut < lastAzimut {
			azimutOffset += 2 * math.Pi
		}
		lastAzimut = currAzimut

		azimut[i] = azimutOffset + currAzimut
		altitude[i] = sunangel.AltitudeSunAngleAt(currTime, horizon.Place)

		currTime = currTime.Add(timeStep)
	}

	altitudeAtAzimut := func(azimutVal float64) float64 {
		leftIndex := 0
		for leftIndex < len(azimut)-2 && azimut[leftIndex] < azimutVal {
			leftIndex++
		}

		return altitude[leftIndex]
	}
	plot.AddFunc2d("Sun", "lines", azimut, altitudeAtAzimut)

	horizonAltitudeAtAzimut := func(azimutVal float64) float64 {
		return horizon.GetAltitude(angle.NormalizeRadians(azimutVal))
	}
	plot.AddFunc2d("Horizon", "lines", azimut, horizonAltitudeAtAzimut)

	plot.SavePlot(fmt.Sprintf("Img/Horizons/%s.png", horizon.Place.Name))
}
