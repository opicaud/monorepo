load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("@bazel_tools//tools/build_defs/repo:git.bzl","git_repository")

def repos():
    maybe(
        http_file,
        name = "pact_prototuf_plugin_archive",
        sha256 = "e3b09357c4ca793b7f0d78716ffe18916a7e72ed346ca549dfed79a4ff85cfc3",
        urls = ["https://github.com/pactflow/pact-protobuf-plugin/releases/download/v-0.3.5/pact-protobuf-plugin-osx-x86_64.gz"],
    )

    maybe(
        http_file,
        name = "pact_plugin_json_archive",
        sha256 = "70fa091ec6728d0077470d7ab1125be02b9b8211b73a552ea37f14e0276b7a52",
        urls = ["https://github.com/pactflow/pact-protobuf-plugin/releases/download/v-0.3.5/pact-plugin.json"],
    )

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


