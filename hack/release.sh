#!/bin/sh

runfiles_dir=$PWD
path=$(dirname $1)
echo "---Release of "$(dirname $1)"---"

if [ -d "monorepo" ]
then
    echo "--> monorepo present, delete it"
    rm -rf monorepo
fi

echo "--> cloning monorepo"

if [ -z "${GH_TOKEN}" ]
then
  echo "--> GH_TOKEN not found, exiting now"
  exit 1
else
  echo "--> GH_TOKEN found, continuing"
  cd $path
  GH_TOKEN=${GH_TOKEN} $runfiles_dir/hack/semantic_release_binary.sh --dry-run
fi