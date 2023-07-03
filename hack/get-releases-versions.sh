#!/bin/sh

getVersionOfApp () {
  firstapp=$(echo "$1" | head -n1)
  semver=$(echo "$firstapp" | awk -F  '-v' '/1/ {print $2}')
  version=v$semver
  app=$(echo "$firstapp"| awk -F  '-v' '/1/ {print $1}')
  echo "STABLE_$(echo "$app" | awk '{ print toupper($0) }' | sed 's/-/_/g') $version"
  echo "STABLE_$(echo "$app" | awk '{ print toupper($0) }'| sed 's/-/_/g')_SEMVER $semver"
  tags=$(echo "$tags" | grep -v "$app")
  nb=$(echo "$tags" | wc -l)
  if [ "$nb" -gt "1" ]; then
      getVersionOfApp "$tags"
  fi

}
git fetch
tags=$(git ls-remote --tags --sort=-committerdate | awk '{ print $2; }' | awk -F  '/' '/1/ {print $3}')
getVersionOfApp "$tags"

apps=$(find -L bazel-bin -name next-version-to-release)
for app in $apps
do
 nextReleaseAndVersion=$(cat "$app" | tail -n 1)
 next=$(echo "$nextReleaseAndVersion" | cut -d " " -f 1)
 version=$(echo "$nextReleaseAndVersion" | cut -d " " -f 2)
 echo "STABLE_$(echo "$next" | awk '{ print toupper($0) }' | sed 's/-/_/g')_NEXT_RELEASE_VERSION $version"
done

echo "STABLE_GH_TOKEN" ${GH_TOKEN}
exit 0