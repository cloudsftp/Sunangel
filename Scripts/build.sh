#!/bin/bash

OUT="Build"
mkdir -p "$OUT"

go build -o "$OUT" github.com/cloudsftp/Sunangel/cmd/sunset
[ $? -gt 0 ] && exit $?

go build -o "$OUT" github.com/cloudsftp/Sunangel/cmd/location
exit $?
