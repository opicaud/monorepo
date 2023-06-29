load("@npm//:defs.bzl", "npm_link_all_packages")
load("@npm__semantic-release__20.1.3//hack:package_json.bzl", "semantic_release_binary")

exports_files([
    "release.sh",
    "find-what-packages-to-release.sh",
    "find-what-packages-to-release-local.sh",
    "release-packages.sh",
    "upgrade_charts.sh"
])
npm_link_all_packages(name = "node_modules")

semantic_release_binary(
    name = "semantic_release_binary",
    data = [
        ":node_modules/@octokit/request-error",
        ":node_modules/@semantic-release/commit-analyzer",
        ":node_modules/@semantic-release/git",
        ":node_modules/@semantic-release/github",
        ":node_modules/@semantic-release/release-notes-generator",
        ":node_modules/semantic-release-monorepo",
    ],
    visibility = ["//visibility:public"],
)

genrule(
    name = "find-what-packages-to-release",
    srcs = [],
    outs = ["packages-to-release"],
    cmd = "./$(location //hack:find-what-packages-to-release.sh) > \"$@\"",
    #tags = ["no-cache"],
    tools = [
        "//hack:find-what-packages-to-release.sh",
        "//hack:find-what-packages-to-release-local.sh",
    ],
    visibility = ["//visibility:private"],
)

load(":release.bzl", "release_me")

release_me()
