#!/bin/sh

toRelease=""
branch="origin/main"
tags=$(git tag --sort=committerdate)
lastTag=$(echo "$tags" | tail -n 1)
lastTagRef=$lastTag

if [ -d "monorepo" ]
then
    echo "--> monorepo present, delete it"
    rm -rf monorepo
fi

echo "--> cloning monorepo"
git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
cd monorepo

changes=$(git diff --name-only "$lastTagRef" "$branch")
if [ $? -ne 0 ]
  then
    echo "--> issues occurred during git diff, exiting now"
    exit 1
fi
echo "--> identify what to release between latest tag $lastTag and $branch.."
for file in $changes; do
  queried=$(bazel query --keep_going --noshow_progress "$file" 2>/dev/null)
  if [ $? -eq 0 ]
    then
       package=$(echo "$queried" | cut -d ':' -f 1)
       hasBeenIdentified=$(echo "$toRelease" | grep "$package")
       if [ "$hasBeenIdentified" = "" ] && [ "$package" != '//' ]
         then
           releaseTarget=$(bazel query --keep_going --noshow_progress "filter("release_me", kind("sh_binary", $package/...))")
           echo "--> $package will be released"
           toRelease="$toRelease $releaseTarget"
       fi
       hasBeenIdentified=""
  fi
done

echo "--> start to effectively release monorepo's components.."
for i in $toRelease
do
  bazel run --noshow_progress "$i"
done