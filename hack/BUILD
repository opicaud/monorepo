load("@npm//:defs.bzl", "npm_link_all_packages")
load("@npm__semantic-release__20.1.3//hack:package_json.bzl", "semantic_release_binary")

exports_files([
    "release.sh",
    "find-what-next-releases-are.sh",
    "get-latest-tags.sh",
    "do-i-need-to-be-released.sh"
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
        ":node_modules/@semantic-release/exec"
    ],
    visibility = ["//visibility:public"],
)

genrule(
    name = "get-latest-tags",
    srcs = [],
    outs = ["latest-tags"],
    cmd = "./$(location //hack:get-latest-tags.sh) --only-apps=1 > \"$@\"",
    tools = [
        "//hack:get-latest-tags.sh",
    ],
    visibility = ["//visibility:public"],
)

load(":release.bzl", "release_me")

release_me()
