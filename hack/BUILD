load("@npm//:defs.bzl", "npm_link_all_packages")
load("@npm__semantic-release__23.0.5//hack:package_json.bzl", "semantic_release_binary")

exports_files([
    "release.sh",
    "find-what-next-releases-are.sh",
    "do-i-need-to-be-released.sh",
    "prepare_release.sh",
])

npm_link_all_packages(name = "node_modules")

semantic_release_binary(
    name = "semantic_release_binary",
    data = [
        ":node_modules/@lomray/semantic-release-monorepo",
        ":node_modules/@octokit/request-error",
        ":node_modules/@semantic-release/commit-analyzer",
        ":node_modules/@semantic-release/exec",
        ":node_modules/@semantic-release/git",
        ":node_modules/@semantic-release/github",
        ":node_modules/@semantic-release/release-notes-generator",
    ],
    visibility = ["//visibility:public"],
)

load(":release.bzl", "release_me")

release_me()
