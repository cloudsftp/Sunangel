# Sunangel ![version](https://img.shields.io/badge/v0.0.0-blue.svg)

[![unit test status](https://github.com/cloudsftp/Sunangel/actions/workflows/core_unit_tests.yaml/badge.svg?branch=develop)](https://github.com/cloudsftp/Sunangel/actions/workflows/core_unit_tests.yaml)
[![unit test status](https://github.com/cloudsftp/Sunangel/actions/workflows/cli_unit_tests.yaml/badge.svg?branch=develop)](https://github.com/cloudsftp/Sunangel/actions/workflows/cli_unit_tests.yaml)
[![build status](https://github.com/cloudsftp/Sunangel/actions/workflows/build.yaml/badge.svg?branch=develop)](https://github.com/cloudsftp/Sunangel/actions/workflows/build.yaml)

## Motivation

Weather Apps and other Sources of information give bad results for the time of sunset.
Especially in locations with many hills.

This project aims to provide a better prediction of the time, the sun actually crosses the horizon.
In order to do so, we use elevation data to compute the horizon profile for a given location.
Then we search for the time, the altitude of the sun is right at the horizon.

## Installation

### From source

```
git clone https://github.com/cloudsftp/Sunangel.git
cd Sunangel
git checkout master # make sure you are at the latest release, not on develop
./install.sh
```

### Binary

Download binaries from the [releases](https://github.com/cloudsftp/Sunangel/releases).

## Usage

### Use `location` to manage stored locations

```
Usage: location COMMAND [OPTIONS]

  COMMAND OPTIONS
  help                     Prints this information
  list                     Lists all stored locations
  add     NAME LAT LONG    Adds a location
  delete  NAME             Deletes a location
```

### Use `sunset` to calculate the sunset for a given location

```
Usage: sunset (help | NAME | LAT LONG) [d=DAYOFFSET] [r=STARTRADIUS]

  Either NAME or LAT LONG is required as the first set of arguments
    If a NAME is entered, the program checks the database for stored locations
    If LAT and LONG are entered, the program uses these coordinates

  DAYOFFSET: Integer offset of the day relativeto today (tomorrow is d=1)
  STARTRADIUS: Integer radius to ignore when computing the horizon (one kilometer is r=1000)
```
