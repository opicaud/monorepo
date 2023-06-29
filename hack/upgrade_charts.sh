#!/bin/sh
chartsToUpgrade=$(bazel query --keep_going --noshow_progress "filter("upgrade_chart", kind("write_source_file", //apps/...))" 2>/dev/null)
for chartToUpgrade in $chartsToUpgrade
do
  bazel run $chartToUpgrade
done
