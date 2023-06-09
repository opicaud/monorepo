#!/bin/sh

if [ -d "monorepo" ]
then
    rm -rf monorepo
fi

git clone --single-branch --branch main --quiet https://github.com/opicaud/monorepo.git
cd monorepo
ls
../hack/find-what-packages-to-release-local.sh "origin/main"