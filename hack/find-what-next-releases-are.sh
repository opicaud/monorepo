#!/bin/sh
runfiles_dir=$PWD
export BAZEL_BINDIR=.
path=$(dirname $3)
if [ -z "${GH_TOKEN}" ]
then
  echo "--> GH_TOKEN not found, exiting now"
  exit 1
fi

if [ -d "monorepo" ]
then
    rm -rf monorepo
fi

git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
cd monorepo/"$path" || exit 1

GH_TOKEN=${GH_TOKEN} $runfiles_dir/$1 --dry-run
version="no"
if [ -f next_release_version ]
then
  version=$(cat next_release_version)
  rm next_release_version
fi
echo $(cat $runfiles_dir/$2 | sed 's/\"//g') $version


