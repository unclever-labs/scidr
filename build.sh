#!/usr/bin/env bash

set -e

rm -rf bin ||:
mkdir  bin ||:

PROJECT=$(basename $(git rev-parse --show-toplevel))
VERSION=$(cat version.txt)
export CGO_ENABLED=0

for arch in amd64 386; do
  export GOARCH=$arch
  for os in darwin linux windows; do
    export GOOS=$os
    go build -o "bin/${PROJECT}-${VERSION}-${GOOS}-${GOARCH}"
    echo "bin/${PROJECT}-${VERSION}-${GOOS}-${GOARCH}"
    shasum -a 256 < "bin/${PROJECT}-${VERSION}-${GOOS}-${GOARCH}" | cut -d ' ' -f1
    echo
  done
done
