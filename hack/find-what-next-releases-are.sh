#!/bin/sh
runfiles_dir=$PWD
export BAZEL_BINDIR=.
path=$(dirname $3)

if [ -d "monorepo" ]
then
    rm -rf monorepo
fi

git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
cd monorepo
cd $path || exit 1

$runfiles_dir/$1 --dry-run 1>/dev/null || exit 1
echo "app=$(cat $runfiles_dir/$2 | sed 's/\"//g')"

has_been_already_released=$(cat current_release_version)
if [ "$has_been_already_released" = "" ]
then
  echo "currentVersion=not-available"
  echo "nextVersion=1.0.0"
  exit 0
else
  echo "currentVersion=$(cat current_release_version)"
fi
if [ -f next_release_version ]
then
  echo "nextVersion=$(cat next_release_version)"
else
  echo "nextVersion=$(cat current_release_version)"
fi
