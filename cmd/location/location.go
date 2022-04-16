package main

import (
	"fmt"
	"os"

	"github.com/cloudsftp/Sunangel/args"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/persist"
)

func main() {
	arguments, err := args.ParseLocationArguments(os.Args)
	if err != nil {
		args.PrintLocationUsage(os.Args[0], err)
	}

	switch arguments.Mode {
	case args.List:
		locations := persist.GetLocations()

		printLocationsList(locations)
	case args.Add:
		place := location.NewLocation(arguments.Latitude, arguments.Longitude)
		persist.AddLocation(arguments.Name, place)

		printLocationsAdded(map[string]*location.Location{arguments.Name: place})
	case args.Delete:
		persist.DeleteLocation(arguments.Name)

		printLocationsDeleted(arguments.Name)
	default:
		args.PrintLocationUsage(os.Args[0], fmt.Errorf("unrecognized command"))
	}
}

func printLocationsList(locations map[string]*location.Location) {
	fmt.Print("\nLocations:\n\n")
	printLocations(locations)
}

func printLocationsAdded(locations map[string]*location.Location) {
	fmt.Print("\nLocations added:\n\n")
	printLocations(locations)
}

func printLocations(locations map[string]*location.Location) {
	for name, place := range locations {
		fmt.Printf("  %-20s  %11.8f, %11.8f\n", name, place.Latitude, place.Longitude)
	}
	fmt.Println()
}

func printLocationsDeleted(name string) {
	fmt.Printf("\nLocations with name %s deleted\n", name)
}
