package args

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cloudsftp/Sunangel/location"
)

type SunArguments struct {
	Place        *location.Location
	RadiusIgnore float64
	DayOffset    int
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

	return &SunArguments{
		Place:        location.NewLocation(latitude, longitude),
		RadiusIgnore: 0,
		DayOffset:    0,
	}, nil
}

func PrintSunUsage() {
	fmt.Printf("Usage: \n")

	os.Exit(2)
}
