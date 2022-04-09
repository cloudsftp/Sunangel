package args

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cloudsftp/Sunangel/location"
)

type SunArguments struct {
	Place       *location.Location
	StartRadius int
	DayOffset   int
}

func NewSunArguments(latitude, longitude float64, startRadius, dayOffset int) *SunArguments {
	return &SunArguments{
		Place: location.NewLocation(
			latitude,
			longitude,
		),
		StartRadius: startRadius,
		DayOffset:   dayOffset,
	}
}

func ParseSunArguments(args []string) (*SunArguments, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("too few arguments")
	}

	latitude, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return nil, fmt.Errorf("first argument is NaN")
	}

	longitude, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return nil, fmt.Errorf("second argument is NaN")
	}

	startRadius := 0
	dayOffset := 0

	for i := 3; i < len(args); i++ {
		parts := strings.Split(args[i], "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("wrong format %s", args[i])
		}

		val, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("could not parse value of %s", args[i])
		}

		switch parts[0] {
		case "r":
			startRadius = val
		case "d":
			dayOffset = val
		default:
			return nil, fmt.Errorf("unrecognized variable %s", args[i])
		}
	}

	return &SunArguments{
		Place:       location.NewLocation(latitude, longitude),
		StartRadius: startRadius,
		DayOffset:   dayOffset,
	}, nil
}

func PrintSunUsage() {
	fmt.Printf("Usage: \n")

	os.Exit(2)
}
