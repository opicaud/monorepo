#!/usr/bin/env sh
bazel run //:gazelle -- update-repos
bazel run //:gazelle -- update
bazel test --test_tag_filters="-integration" //...
