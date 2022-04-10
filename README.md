# Sunangel ![version](https://img.shields.io/badge/v0.0.0-blue.svg)

![unit test status](https://github.com/cloudsftp/Sunangel/actions/workflows/core_unit_tests.yaml/badge.svg?branch=develop)
![unit test status](https://github.com/cloudsftp/Sunangel/actions/workflows/cli_unit_tests.yaml/badge.svg?branch=develop)

## Motivation

Weather Apps and other Sources of information give bad results for the time of sunset.
Especially in locations with many hills.

This project aims to provide a better prediction of the time, the sun actually crosses the horizon.
In order to do so, we use elevation data to compute the horizon profile for a given location.
Then we search for the time, the altitude of the sun is right at the horizon.
