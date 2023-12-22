#!/bin/sh

runfiles_dir=$PWD
export BAZEL_BINDIR=.
path=$(dirname $1)

echo "---Release of "$(dirname $1)"---"

if [ -z "${GH_TOKEN}" ]
then
  echo "--> GH_TOKEN not found, exiting now"
  exit 1
else
  echo "--> GH_TOKEN found, continuing"
  cd $BUILD_WORKSPACE_DIRECTORY/$path || exit 1
  GH_TOKEN=${GH_TOKEN} $runfiles_dir/hack/semantic_release_binary.sh
  rm -f next_release_version
fi