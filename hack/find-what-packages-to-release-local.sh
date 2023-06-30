#!/bin/sh

toRelease=""
branch=${1:-"main"}
tags=$(git tag --sort=committerdate)
lastTag=$(echo "$tags" | tail -n 1)
lastTagRef=$lastTag
changes=$(git diff --name-only "$lastTagRef" "$branch")
if [ $? -ne 0 ]
  then
    exit 1
fi
releaseCandidates=$(bazel query --keep_going --noshow_progress "filter("release_me", kind("sh_binary", //...))")
for releaseCandidate in $releaseCandidates;
do
  rootPackage=$(echo "$releaseCandidate" | cut -d ':' -f 1 | sed 's/\/\///g')
  for change in $changes
  do
    found=$(echo "$change" | grep -c "^$rootPackage")
    hasBeenAlreadyFound=$(echo "$toRelease" | grep "$releaseCandidate")
    if [ "$found" -eq 1 ] && [ "$hasBeenAlreadyFound" == "" ]
     then
        toRelease="$toRelease $releaseCandidate"
     fi
  done
done
echo "$toRelease"