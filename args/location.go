package args

import (
	"fmt"
	"os"
	"strconv"
)

type LocationMode int

const (
	List LocationMode = iota
	Add
	Delete
)

type LocationArguments struct {
	Mode      LocationMode
	Name      string
	Latitude  float64
	Longitude float64
}

func ParseLocationArguments(args []string) (*LocationArguments, error) {
	arguments := &LocationArguments{}
	if len(args) < 2 {
		return nil, fmt.Errorf("too few argumeents")
	}

	switch args[1] {
	case "list":
		arguments.Mode = List

		if len(args) > 2 {
			return nil, fmt.Errorf("too many arguments for command list")
		}
	case "add":
		arguments.Mode = Add
		if len(args) != 5 {
			return nil, fmt.Errorf("wrong number of arguments for command add")
		}

		arguments.Name = args[2]

		var err error
		arguments.Latitude, err = strconv.ParseFloat(args[3], 64)
		if err != nil {
			return nil, fmt.Errorf("third argument is NaN")
		}

		arguments.Longitude, err = strconv.ParseFloat(args[4], 64)
		if err != nil {
			return nil, fmt.Errorf("fourth argument is NaN")
		}
	case "delete":
		arguments.Mode = Delete
		if len(args) != 3 {
			return nil, fmt.Errorf("wrong number of arsuments for command delete")
		}

		arguments.Name = args[2]
	case "help":
		return nil, fmt.Errorf("help requested")
	default:
		return nil, fmt.Errorf("%s is not recognized as command", args[1])
	}

	return arguments, nil
}

func PrintLocationUsage(cmd string, err error) {
	fmt.Printf("%v\n\n", err)
	fmt.Printf("Usage: %s COMMAND [OPTIONS]\n\n", cmd)
	fmt.Println("  COMMAND OPTIONS")
	fmt.Println("  help                     Prints this information")
	fmt.Println("  list                     Lists all stored locations")
	fmt.Println("  add     NAME LAT LONG    Adds a location")
	fmt.Println("  delete  NAME             Deletes a location")
	fmt.Println()

	os.Exit(2)
}
