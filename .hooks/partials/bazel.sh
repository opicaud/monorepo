#!/usr/bin/env sh
bazel run //:gazelle
bazel test --test_tag_filters="-integration" //...
