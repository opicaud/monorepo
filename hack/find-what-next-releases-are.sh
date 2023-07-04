#!/bin/sh
runfiles_dir=$PWD
export BAZEL_BINDIR=.
path=$(dirname $3)
version="no"

if [ -z "${GH_TOKEN}" ] || [ "${GH_TOKEN}" = "default" ]
then
  echo "$path" $version > $runfiles_dir/$OUT
  exit 0
fi

if [ -d "monorepo" ]
then
    rm -rf monorepo
fi

git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
cd monorepo/"$path" || exit 1

$runfiles_dir/$1 --dry-run || exit 1

if [ -f next_release_version ]
then
  version=$(cat next_release_version)
  rm next_release_version
fi

echo $(cat $runfiles_dir/$2 | sed 's/\"//g') "$version" > $runfiles_dir/$OUT
