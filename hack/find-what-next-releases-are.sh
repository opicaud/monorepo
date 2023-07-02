#!/bin/sh
runfiles_dir=$PWD
export BAZEL_BINDIR=.

if [ -d "monorepo" ]
then
    rm -rf monorepo
fi

git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
cd monorepo/hack || exit 1

GH_TOKEN=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 $runfiles_dir/$1 --dry-run
version="no"
if [ -f next_release_version ]
then
  version=$(cat next_release_version)
  rm next_release_version
fi
echo $(cat $runfiles_dir/$2 | sed 's/\"//g') $version


