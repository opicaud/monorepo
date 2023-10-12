load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("@bazel_tools//tools/build_defs/repo:git.bzl","git_repository")

def repos():

    maybe(
        http_file,
        name = "pact_verifier_cli_archive",
        sha256 = "77ffc38f4564cfef42f64b9eb33bebfc4d787e65ef7ff7213640a3d63d2cf5a7",
        urls = ["https://github.com/pact-foundation/pact-reference/releases/download/pact_verifier_cli-v1.0.1/pact_verifier_cli-osx-x86_64.gz"],
    )

    maybe(
        http_file,
        name = "pact_ffi_archive",
        sha256 ="b8c87e2cc2f83ae9e79678d3288f2f9f7cea27d023576f565d8a203441600a59",
        urls = ["https://github.com/pact-foundation/pact-reference/releases/download/libpact_ffi-v0.4.9/libpact_ffi-osx-x86_64.dylib.gz"]
    )
