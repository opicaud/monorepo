#!/bin/sh
runfiles_dir=$PWD
export BAZEL_BINDIR=.
path=$(dirname $3)

if [ -z "${GH_TOKEN}" ] || [ "${GH_TOKEN}" = "default" ]
then
  echo "app=$(cat $runfiles_dir/$2 | sed 's/\"//g')"
  echo "currentVersion=not-available"
  echo "nextVersion=not-available"
  exit 0
fi

if [ -d "monorepo" ]
then
    rm -rf monorepo
fi

git clone --single-branch --branch test --quiet https://github.com/opicaud/monorepo.git
cd monorepo
cd $path || exit 1

$runfiles_dir/$1 --dry-run 1>/dev/null || exit 1
echo "app=$(cat $runfiles_dir/$2 | sed 's/\"//g')"

if [ -f current_release_version ]
then
  echo "currentVersion=$(cat current_release_version)"
else
  echo "currentVersion=not-available"
  echo "nextVersion=1.0.0"
  exit 0
fi
if [ -f next_release_version ]
then
  echo "nextVersion=$(cat next_release_version)"
else
  echo "nextVersion=$(cat current_release_version)"
fi
