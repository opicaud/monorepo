#!/bin/sh
runfiles_dir=$PWD
export BAZEL_BINDIR=.
path=$(dirname $3)

if [ -z "${GH_TOKEN}" ] || [ "${GH_TOKEN}" = "default" ]
then
  echo "app=$(cat $runfiles_dir/$2 | sed 's/\"//g')" > $OUT
  echo "currentVersion=not-available" >> $OUT
  echo "nextVersion=not-available" >> $OUT
  exit 0
fi

if [ -d "monorepo" ]
then
    rm -rf monorepo
fi

git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
cd monorepo
releaserc=$(cat << \EOF
{
  "branches": ["main"],
  "extends": "semantic-release-monorepo",
  "plugins": [
    "@semantic-release/commit-analyzer",
    "@semantic-release/release-notes-generator",
    "@semantic-release/git",
    "@semantic-release/github",
    ["@semantic-release/exec", {
        "analyzeCommitsCmd": "echo ${lastRelease.version} > current_release_version",
        "verifyReleaseCmd": "echo ${nextRelease.version} > next_release_version",
    }]
  ]
}
EOF
)
echo "$releaserc" > .releaserc
cd $path || exit 1

$runfiles_dir/$1 --dry-run || exit 1
echo "app=$(cat $runfiles_dir/$2 | sed 's/\"//g')" > $runfiles_dir/$OUT
echo "currentVersion=$(cat current_release_version)" >> $runfiles_dir/$OUT
if [ -f next_release_version ]
then
  echo "nextVersion=$(cat next_release_version)" >> $runfiles_dir/$OUT
else
  echo "nextVersion=$(cat current_release_version)" >> $runfiles_dir/$OUT
fi
