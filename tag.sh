#!/usr/bin/env bash

VERSION=$(cat version.txt)

echo "Tagging: $VERSION"
git tag $VERSION

echo "git push origin $VERSION"
git push origin $VERSION
