#!/bin/sh

component="STABLE_$(cat "$2" | sed 's/\"//g' | awk '{ print toupper($0) }' | sed 's/-/_/g')"
currentVersion=$(cat $1 | grep $component | grep "SEMVER" | cut -d " " -f 2)
nextVersion=$(cat "$3" | tail -n 1 | cut -d " " -f 2)
if [ "$nextVersion" != "$currentVersion" ]
then
  echo "YES"
else
  echo "NO"
fi