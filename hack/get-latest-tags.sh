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
if [ -d "monorepo" ]
then
    rm -rf monorepo
fi
git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
cd monorepo/ || exit 1
tags=$(git ls-remote --tags --sort=-committerdate | awk '{ print $2; }' | awk -F  '/' '/1/ {print $3}')
getVersionOfApp "$tags"