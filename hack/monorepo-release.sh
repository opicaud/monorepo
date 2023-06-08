#!/bin/sh

toRelease=""
branch="origin/main"
tags=$(git tag --sort=committerdate)
lastTag=$(echo "$tags" | tail -n 1)
lastTagRef="refs/tags/"$lastTag

cd $BUILD_WORKSPACE_DIRECTORY
echo "Identify what to release between latest tag $lastTag and $branch.."
for file in $(git diff --name-only "$lastTagRef" "$branch" ); do
  queried=$(bazel query --keep_going --noshow_progress "$file" 2>/dev/null)
  if [ $? -eq 0 ]
    then
       package=$(echo "$queried" | cut -d ':' -f 1)
       hasBeenIdentified=$(echo "$toRelease" | grep "$package")
       if [ "$hasBeenIdentified" = "" ] && [ "$package" != '//' ]
         then
           releaseTarget=$(bazel query --keep_going --noshow_progress "filter("release_me", kind("sh_binary", $package/...))")
           echo "$package will be released"
           toRelease="$toRelease $releaseTarget"
       fi
       hasBeenIdentified=""
  fi
done

echo "Start to effectively release monorepo's components.."
for i in $toRelease
do
  bazel run --noshow_progress "$i"
done