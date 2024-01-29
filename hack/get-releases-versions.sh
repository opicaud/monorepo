#!/bin/sh

apps=$(find -L bazel-bin -name next-version-to-release)
for appFeature in $apps
do
 . $appFeature
 echo "STABLE_$(echo "$app" | awk '{ print toupper($0) }' | sed 's/-/_/g')_NEXT_RELEASE_VERSION v$nextVersion"
 echo "STABLE_$(echo "$app" | awk '{ print toupper($0) }' | sed 's/-/_/g')_NEXT_RELEASE_SEMVER $nextVersion"

done

exit 0