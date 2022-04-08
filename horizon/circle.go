package horizon

import "math"

type offsetCoordinates struct {
	latitude  int
	longitude int
}

// circle computes a circle.
// It starts ar (0, r) and then moves clockwise
func circle(r int) []offsetCoordinates {
	arc := arc(r)

	quadrantSize := 2*len(arc) - 1
	ignoreLastPixelHalfOfTheTime := false
	if arc[len(arc)-1].latitude == arc[len(arc)-1].longitude {
		ignoreLastPixelHalfOfTheTime = true
		quadrantSize -= 1
	}
	circleSize := 4 * quadrantSize
	circle := make([]offsetCoordinates, circleSize)

	for i := 0; i < len(arc); i++ {
		x := arc[i].latitude
		y := arc[i].longitude

		circle[i] = offsetCoordinates{latitude: x, longitude: y}
		circle[quadrantSize+i] = offsetCoordinates{latitude: -y, longitude: x}
		circle[2*quadrantSize+i] = offsetCoordinates{latitude: -x, longitude: -y}
		circle[3*quadrantSize+i] = offsetCoordinates{latitude: y, longitude: -x}

		if i > 0 && (i < len(arc)-1 || !ignoreLastPixelHalfOfTheTime) {
			circle[quadrantSize-i] = offsetCoordinates{latitude: y, longitude: x}
			circle[2*quadrantSize-i] = offsetCoordinates{latitude: -x, longitude: y}
			circle[3*quadrantSize-i] = offsetCoordinates{latitude: -y, longitude: -x}
			circle[4*quadrantSize-i] = offsetCoordinates{latitude: x, longitude: -y}
		}
	}

	return circle
}

// arc computes an eightth circle.
// It starts at (r, 0) and then moves counter clockwise.
// Note: the last coordinates could be x = y

func arc(r int) []offsetCoordinates {
	arc := make([]offsetCoordinates, 1)
	x := r
	y := 0

	arc[0] = offsetCoordinates{latitude: x, longitude: y}

	for x > y {
		if radiusError(x-1, y+1, r) < radiusError(x, y+1, r) {
			x -= 1
		}
		y += 1

		if x < y {
			break
		}

		arc = append(arc, offsetCoordinates{latitude: x, longitude: y})
	}

	return arc
}

func radiusError(x, y, r int) int {
	return int(math.Abs(float64(x*x + y*y - r*r)))
}
