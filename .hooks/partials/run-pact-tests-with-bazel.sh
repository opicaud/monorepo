#!/usr/bin/env sh
targets=$(bazel query "attr(tags, '\\bpact_test\\b', //...)")
for target in $targets
do
  echo "Running target: ""$target"
  bazel run "$target"
done
