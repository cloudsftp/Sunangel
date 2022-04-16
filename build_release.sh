#!/bin/bash

OUT="Build"
ARCHIVE_NAME="SunangelBinaries"

./build.sh

cd "$OUT"
tar cfz "${ARCHIVE_NAME}Linux.tar.gz" location sunset
rm location sunset
cd -

GOOS="darwin" ./build.sh

cd "$OUT"
zip "${ARCHIVE_NAME}Darwin.zip" location sunset
rm location sunset
cd -

GOOS="windows" ./build.sh

cd "$OUT"
zip "${ARCHIVE_NAME}Windows.zip" location.exe sunset.exe
rm location.exe sunset.exe
cd -

mv "Build/${ARCHIVE_NAME}"* .
