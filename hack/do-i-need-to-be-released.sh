#!/bin/sh

. "$1"
if [ "$nextVersion" != "$currentVersion" ]
then
  echo "YES"
else
  echo "NO"
fi