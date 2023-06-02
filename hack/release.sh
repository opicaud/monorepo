#!/bin/sh

runfiles_dir=$PWD
path=$(dirname $1)
echo "---Release of "$(dirname $1)"---"

if [ -d "monorepo" ]
then
    echo "--> monorepo exists in cache, rebase it"
    cd monorepo
    git fetch --quiet
    git rebase origin/main --quiet
else
    echo "--> monorepo does not exist in cache, cloning it"
    git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
    cd monorepo
fi

if [ -z "${GH_TOKEN}" ]
then
  echo "--> GH_TOKEN not found, exiting now"
  exit 1
else
  echo "--> GH_TOKEN found, continuing"
  cd $path
  GH_TOKEN=${GH_TOKEN} $runfiles_dir/hack/semantic_release_binary.sh --dry-run
fi