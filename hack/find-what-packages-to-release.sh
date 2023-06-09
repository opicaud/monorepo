#!/bin/sh

if [ -d "monorepo" ]
then
    rm -rf monorepo
fi

git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
cd monorepo

toRelease=""
branch="origin/main"
tags=$(git tag --sort=committerdate)
lastTag=$(echo "$tags" | tail -n 1)
lastTagRef=$lastTag
changes=$(git diff --name-only "$lastTagRef" "$branch")
if [ $? -ne 0 ]
  then
    exit 1
fi
for file in $changes; do
  queried=$(bazel query --keep_going --noshow_progress "$file" 2>/dev/null)
  if [ $? -eq 0 ]
    then
       package=$(echo "$queried" | cut -d ':' -f 1)
       hasBeenIdentified=$(echo "$toRelease" | grep "$package")
       if [ "$hasBeenIdentified" = "" ] && [ "$package" != '//' ]
         then
           releaseTarget=$(bazel query --keep_going --noshow_progress "filter("release_me", kind("sh_binary", $package/...))" 2>/dev/null)
           if [ "$releaseTarget" != "" ]
           then
              toRelease="$toRelease $package"
           fi
       fi
       hasBeenIdentified=""
  fi
done

echo "$toRelease"