#!/usr/bin/env bash

VERSION=$(cat version.txt)

echo "Tagging: $VERSION"
git tag $VERSION
