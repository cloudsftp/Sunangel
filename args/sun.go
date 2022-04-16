package args

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cloudsftp/Sunangel/location"
)

type SunMode int

const (
	Coordinates SunMode = iota
	Name
)

type SunArguments struct {
	Mode        SunMode
	Name        string
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
	if len(args) < 2 {
		return nil, fmt.Errorf("too few arguments")
	}

	arguments := &SunArguments{Mode: Coordinates}
	latitude, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		arguments.Mode = Name
	}

	var variablesStartIndex int
	switch arguments.Mode {
	case Coordinates:
		if len(args) < 3 {
			return nil, fmt.Errorf("too few arguments for coordinates mode")
		}

		longitude, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return nil, fmt.Errorf("second argument is NaN")
		}

		arguments.Place = location.NewLocation(latitude, longitude)
		variablesStartIndex = 3
	case Name:
		arguments.Name = args[1]
		variablesStartIndex = 2

		if arguments.Name == "help" {
			return nil, fmt.Errorf("help requested")
		}
	default:
		return nil, fmt.Errorf("unrecognized mode")
	}

	for i := variablesStartIndex; i < len(args); i++ {
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
			arguments.StartRadius = val
		case "d":
			arguments.DayOffset = val
		default:
			return nil, fmt.Errorf("unrecognized variable %s", args[i])
		}
	}

	return arguments, nil
}

func PrintSunUsage(cmd string, err error) {
	fmt.Printf("%v\n\n", err)
	fmt.Printf("Usage: %s (help | NAME | LAT LONG) [d=DAYOFFSET] [r=STARTRADIUS]\n\n", cmd)
	fmt.Println("  Either NAME or LAT LONG is required as the first set of arguments")
	fmt.Println("    If a NAME is entered, the program checks the database for stored locations")
	fmt.Println("    If LAT and LONG are entered, the program uses these coordinates")
	fmt.Println()
	fmt.Println("  DAYOFFSET: Integer offset of the day relativeto today (tomorrow is d=1)")
	fmt.Println("  STARTRADIUS: Integer radius to ignore when computing the horizon (one kilometer is r=1000)")
	fmt.Println()

	os.Exit(2)
}
