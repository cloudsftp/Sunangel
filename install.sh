#!/bin/bash

go install github.com/cloudsftp/Sunangel/cmd/sunset
[ $? -gt 0 ] && exit $?

go install github.com/cloudsftp/Sunangel/cmd/location
exit $?
