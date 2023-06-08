#!/bin/sh

if [ -d "monorepo" ]
then
    echo "--> monorepo present, delete it"
    rm -rf monorepo
fi

echo "--> cloning monorepo"
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
           releaseTarget=$(bazel query --keep_going --noshow_progress "filter("release_me", kind("sh_binary", $package/...))" 2>/dev/null)
           if [ "$releaseTarget" != "" ]
           then
              echo "--> $package will be released"
              toRelease="$toRelease $package"
           else
              echo "--> WARN: not release target found for $package"
           fi
       fi
       hasBeenIdentified=""
  fi
done

echo "$toRelease" > packages-to-release
echo "Release file is available here: $PWD/packages-to-release"